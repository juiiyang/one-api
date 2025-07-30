package model

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Laisky/errors/v2"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/songquanpeng/one-api/common"
	"github.com/songquanpeng/one-api/common/config"
	"github.com/songquanpeng/one-api/common/env"
	"github.com/songquanpeng/one-api/common/helper"
	"github.com/songquanpeng/one-api/common/logger"
	"github.com/songquanpeng/one-api/common/random"
	// glogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var LOG_DB *gorm.DB

func CreateRootAccountIfNeed() error {
	var user User
	//if user.Status != util.UserStatusEnabled {
	if err := DB.First(&user).Error; err != nil {
		logger.Logger.Info("no user exists, creating a root user for you: username is root, password is 123456")
		hashedPassword, err := common.Password2Hash("123456")
		if err != nil {
			return errors.WithStack(err)
		}
		accessToken := random.GetUUID()
		if config.InitialRootAccessToken != "" {
			accessToken = config.InitialRootAccessToken
		}
		rootUser := User{
			Username:    "root",
			Password:    hashedPassword,
			Role:        RoleRootUser,
			Status:      UserStatusEnabled,
			DisplayName: "Root User",
			AccessToken: accessToken,
			Quota:       500000000000000,
		}
		DB.Create(&rootUser)
		if config.InitialRootToken != "" {
			logger.Logger.Info("creating initial root token as requested")
			token := Token{
				Id:             1,
				UserId:         rootUser.Id,
				Key:            config.InitialRootToken,
				Status:         TokenStatusEnabled,
				Name:           "Initial Root Token",
				CreatedTime:    helper.GetTimestamp(),
				AccessedTime:   helper.GetTimestamp(),
				ExpiredTime:    -1,
				RemainQuota:    500000000000000,
				UnlimitedQuota: true,
			}
			DB.Create(&token)
		}
	}
	return nil
}

func chooseDB(envName string) (*gorm.DB, error) {
	dsn := os.Getenv(envName)

	switch {
	case strings.HasPrefix(dsn, "postgres://"):
		// Use PostgreSQL
		return openPostgreSQL(dsn)
	case dsn != "":
		// Use MySQL
		return openMySQL(dsn)
	default:
		// Use SQLite
		return openSQLite()
	}
}

func openPostgreSQL(dsn string) (*gorm.DB, error) {
	logger.Logger.Info("using PostgreSQL as database")
	common.UsingPostgreSQL = true
	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		PrepareStmt: true, // precompile SQL
		// Logger: glogger.Default.LogMode(glogger.Info),  // debug sql
	})
}

func openMySQL(dsn string) (*gorm.DB, error) {
	logger.Logger.Info("using MySQL as database")
	common.UsingMySQL = true
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}

func openSQLite() (*gorm.DB, error) {
	logger.Logger.Info("SQL_DSN not set, using SQLite as database")
	common.UsingSQLite = true
	dsn := fmt.Sprintf("%s?_busy_timeout=%d", common.SQLitePath, common.SQLiteBusyTimeout)
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true, // precompile SQL
	})
}

func InitDB() {
	var err error
	DB, err = chooseDB("SQL_DSN")
	if err != nil {
		logger.Logger.Fatal("failed to initialize database: " + err.Error())
		return
	}

	if config.DebugSQLEnabled {
		logger.Logger.Debug("debug sql enabled")
		DB = DB.Debug()
	}

	sqlDB := setDBConns(DB)

	if !config.IsMasterNode {
		return
	}

	if common.UsingMySQL {
		_, _ = sqlDB.Exec("DROP INDEX idx_channels_key ON channels;") // TODO: delete this line when most users have upgraded
	}

	logger.Logger.Info("database migration started")
	if err = migrateDB(); err != nil {
		logger.Logger.Fatal("failed to migrate database: " + err.Error())
		return
	}
	logger.Logger.Info("database migrated")

	// Migrate ModelConfigs and ModelMapping columns from varchar(1024) to text
	if err = MigrateChannelFieldsToText(); err != nil {
		logger.Logger.Error("failed to migrate channel field types: " + err.Error())
		// Don't fail startup for this migration, just log the error
	}

	// Migrate existing ModelConfigs data from old format to new format
	if err = MigrateAllChannelModelConfigs(); err != nil {
		logger.Logger.Error("failed to migrate channel ModelConfigs: " + err.Error())
		// Don't fail startup for this migration, just log the error
	}
}

func migrateDB() error {
	var err error
	if err = DB.AutoMigrate(&Channel{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&Token{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&User{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&Option{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&Redemption{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&Ability{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&Log{}); err != nil {
		return err
	}
	if err = DB.AutoMigrate(&UserRequestCost{}); err != nil {
		return err
	}
	return nil
}

func InitLogDB() {
	if os.Getenv("LOG_SQL_DSN") == "" {
		LOG_DB = DB
		return
	}

	logger.Logger.Info("using secondary database for table logs")
	var err error
	LOG_DB, err = chooseDB("LOG_SQL_DSN")
	if err != nil {
		logger.Logger.Fatal("failed to initialize secondary database: " + err.Error())
		return
	}

	setDBConns(LOG_DB)

	if !config.IsMasterNode {
		return
	}

	logger.Logger.Info("secondary database migration started")
	err = migrateLOGDB()
	if err != nil {
		logger.Logger.Fatal("failed to migrate secondary database: " + err.Error())
		return
	}
	logger.Logger.Info("secondary database migrated")
}

func migrateLOGDB() error {
	var err error
	if err = LOG_DB.AutoMigrate(&Log{}); err != nil {
		return err
	}
	return nil
}

func setDBConns(db *gorm.DB) *sql.DB {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Logger.Fatal("failed to connect database: " + err.Error())
		return nil
	}

	// Increase default connection pool sizes to handle billing load better
	maxIdleConns := env.Int("SQL_MAX_IDLE_CONNS", 200)  // Increased from 100
	maxOpenConns := env.Int("SQL_MAX_OPEN_CONNS", 2000) // Increased from 1000
	maxLifetime := env.Int("SQL_MAX_LIFETIME", 300)     // Increased from 60 seconds

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))

	// Log connection pool settings for monitoring
	logger.Logger.Info(fmt.Sprintf("Database connection pool configured: MaxIdle=%d, MaxOpen=%d, MaxLifetime=%ds",
		maxIdleConns, maxOpenConns, maxLifetime))

	// Start connection pool monitoring goroutine
	go monitorDBConnections(sqlDB)

	return sqlDB
}

// monitorDBConnections monitors database connection pool health
func monitorDBConnections(sqlDB *sql.DB) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		stats := sqlDB.Stats()

		// Log warning if connection pool is under stress
		if stats.InUse > int(float64(stats.MaxOpenConnections)*0.8) {
			logger.Logger.Error(fmt.Sprintf("HIGH DB CONNECTION USAGE: InUse=%d/%d (%.1f%%), Idle=%d, WaitCount=%d, WaitDuration=%v",
				stats.InUse, stats.MaxOpenConnections,
				float64(stats.InUse)/float64(stats.MaxOpenConnections)*100,
				stats.Idle, stats.WaitCount, stats.WaitDuration))
		}

		// Log critical error if we're hitting connection limits
		if stats.WaitCount > 0 && stats.WaitDuration > time.Second {
			logger.Logger.Error(fmt.Sprintf("CRITICAL DB CONNECTION BOTTLENECK: WaitCount=%d, WaitDuration=%v - Consider increasing SQL_MAX_OPEN_CONNS",
				stats.WaitCount, stats.WaitDuration))
		}
	}
}

func closeDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return errors.WithStack(err)
	}
	err = sqlDB.Close()
	return errors.WithStack(err)
}

func CloseDB() error {
	if LOG_DB != DB {
		err := closeDB(LOG_DB)
		if err != nil {
			return err
		}
	}
	return closeDB(DB)
}
