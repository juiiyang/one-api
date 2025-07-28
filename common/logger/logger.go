package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	glog "github.com/Laisky/go-utils/v5/log"
	"github.com/gin-gonic/gin"

	"github.com/songquanpeng/one-api/common/config"
)

var (
	Logger       glog.Logger
	setupLogOnce sync.Once
	initLogOnce  sync.Once
)

// init initializes the logger automatically when the package is imported
func init() {
	initLogger()
}

// initLogger initializes the go-utils logger
func initLogger() {
	initLogOnce.Do(func() {
		var err error
		level := glog.LevelInfo
		if config.DebugEnabled {
			level = glog.LevelDebug
		}

		Logger, err = glog.NewConsoleWithName("one-api", level)
		if err != nil {
			panic(fmt.Sprintf("failed to create logger: %+v", err))
		}
	})
}

func SetupLogger() {
	setupLogOnce.Do(func() {
		if LogDir != "" {
			var logPath string
			if config.OnlyOneLogFile {
				logPath = filepath.Join(LogDir, "oneapi.log")
			} else {
				logPath = filepath.Join(LogDir, fmt.Sprintf("oneapi-%s.log", time.Now().Format("20060102")))
			}
			fd, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal("failed to open log file")
			}
			gin.DefaultWriter = io.MultiWriter(os.Stdout, fd)
			gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, fd)
		}
	})
}
