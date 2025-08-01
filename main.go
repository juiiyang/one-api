package main

import (
	"embed"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	gmw "github.com/Laisky/gin-middlewares/v6"
	glog "github.com/Laisky/go-utils/v5/log"
	"github.com/Laisky/zap"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/songquanpeng/one-api/common"
	"github.com/songquanpeng/one-api/common/client"
	"github.com/songquanpeng/one-api/common/config"
	"github.com/songquanpeng/one-api/common/i18n"
	"github.com/songquanpeng/one-api/common/logger"
	"github.com/songquanpeng/one-api/controller"
	"github.com/songquanpeng/one-api/middleware"
	"github.com/songquanpeng/one-api/model"
	"github.com/songquanpeng/one-api/monitor"
	"github.com/songquanpeng/one-api/relay"
	"github.com/songquanpeng/one-api/relay/adaptor/openai"
	"github.com/songquanpeng/one-api/router"
)

//go:embed web/build/*

var buildFS embed.FS

func main() {
	common.Init()
	logger.SetupLogger()
	logger.Logger.Info(fmt.Sprintf("One API %s started", common.Version))

	if os.Getenv("GIN_MODE") != gin.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}
	if config.DebugEnabled {
		logger.Logger.Info("running in debug mode")
	}

	// Initialize SQL Database
	model.InitDB()
	model.InitLogDB()

	var err error
	err = model.CreateRootAccountIfNeed()
	if err != nil {
		logger.Logger.Fatal("database init error", zap.Error(err))
	}
	defer func() {
		err := model.CloseDB()
		if err != nil {
			logger.Logger.Fatal("failed to close database", zap.Error(err))
		}
	}()

	// Initialize Redis
	err = common.InitRedisClient()
	if err != nil {
		logger.Logger.Fatal("failed to initialize Redis: " + err.Error())
	}

	// Initialize options
	model.InitOptionMap()
	logger.Logger.Info(fmt.Sprintf("using theme %s", config.Theme))
	if common.RedisEnabled {
		// for compatibility with old versions
		config.MemoryCacheEnabled = true
	}
	if config.MemoryCacheEnabled {
		logger.Logger.Info("memory cache enabled")
		logger.Logger.Info(fmt.Sprintf("sync frequency: %d seconds", config.SyncFrequency))
		model.InitChannelCache()
	}
	if config.MemoryCacheEnabled {
		go model.SyncOptions(config.SyncFrequency)
		go model.SyncChannelCache(config.SyncFrequency)
	}
	if os.Getenv("CHANNEL_TEST_FREQUENCY") != "" {
		frequency, err := strconv.Atoi(os.Getenv("CHANNEL_TEST_FREQUENCY"))
		if err != nil {
			logger.Logger.Fatal("failed to parse CHANNEL_TEST_FREQUENCY: " + err.Error())
		}
		go controller.AutomaticallyTestChannels(frequency)
	}
	if os.Getenv("BATCH_UPDATE_ENABLED") == "true" {
		config.BatchUpdateEnabled = true
		logger.Logger.Info("batch update enabled with interval " + strconv.Itoa(config.BatchUpdateInterval) + "s")
		model.InitBatchUpdater()
	}
	if config.EnableMetric {
		logger.Logger.Info("metric enabled, will disable channel if too much request failed")
	}

	// Initialize Prometheus monitoring
	if config.EnablePrometheusMetrics {
		startTime := time.Unix(common.StartTime, 0)
		if err := monitor.InitPrometheusMonitoring(common.Version, startTime.Format(time.RFC3339), runtime.Version(), startTime); err != nil {
			logger.Logger.Fatal("failed to initialize Prometheus monitoring: " + err.Error())
		}
		logger.Logger.Info("Prometheus monitoring initialized")

		// Initialize database monitoring
		if err := model.InitPrometheusDBMonitoring(); err != nil {
			logger.Logger.Fatal("failed to initialize database monitoring: " + err.Error())
		}

		// Initialize Redis monitoring if enabled
		if common.RedisEnabled {
			common.InitPrometheusRedisMonitoring()
		}
	}

	openai.InitTokenEncoders()
	client.Init()

	// Initialize global pricing manager
	relay.InitializeGlobalPricing()

	// Initialize i18n
	if err := i18n.Init(); err != nil {
		logger.Logger.Fatal("failed to initialize i18n", zap.Error(err))
	}

	// Ensure logger is initialized
	logger.SetupLogger()

	logLevel := glog.LevelInfo
	if config.DebugEnabled {
		logLevel = glog.LevelDebug
	}

	// Initialize HTTP server
	server := gin.New()
	server.RedirectTrailingSlash = false
	server.Use(
		gin.Recovery(),
		gmw.NewLoggerMiddleware(
			gmw.WithLoggerMwColored(),
			gmw.WithLevel(logLevel.String()),
			gmw.WithLogger(logger.Logger.Named("gin")),
		),
	)
	// This will cause SSE not to work!!!
	//server.Use(gzip.Gzip(gzip.DefaultCompression))
	server.Use(middleware.RequestId())
	server.Use(middleware.Language())

	// Add Prometheus middleware if enabled
	if config.EnablePrometheusMetrics {
		server.Use(middleware.PrometheusMiddleware())
		server.Use(middleware.PrometheusRateLimitMiddleware())
	}

	// middleware.SetUpLogger(server)

	// Initialize session store
	sessionSecret, err := base64.StdEncoding.DecodeString(config.SessionSecret)
	var sessionStore cookie.Store
	if err != nil {
		logger.Logger.Info("session secret is not base64 encoded, using raw value instead")
		sessionStore = cookie.NewStore([]byte(config.SessionSecret))
	} else {
		sessionStore = cookie.NewStore(sessionSecret, sessionSecret)
	}

	if config.DisableCookieSecret {
		logger.Logger.Warn("DISABLE_COOKIE_SECURE is set, using insecure cookie store")
		sessionStore.Options(sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			SameSite: http.SameSiteLaxMode,
			Secure:   false,
		})
	}
	server.Use(sessions.Sessions("session", sessionStore))

	// Add Prometheus metrics endpoint if enabled
	if config.EnablePrometheusMetrics {
		server.GET("/metrics", middleware.AdminAuth(), gin.WrapH(promhttp.Handler()))
		logger.Logger.Info("Prometheus metrics endpoint available at /metrics")
	}

	router.SetRouter(server, buildFS)
	var port = os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(*common.Port)
	}
	logger.Logger.Info(fmt.Sprintf("server started on http://localhost:%s", port))
	err = server.Run(":" + port)
	if err != nil {
		logger.Logger.Fatal("failed to start HTTP server: " + err.Error())
	}
}
