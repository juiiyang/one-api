package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	"github.com/songquanpeng/one-api/common"
	"github.com/songquanpeng/one-api/common/logger"
)

func RelayPanicRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Logger.Error(fmt.Sprintf("panic detected: %v", err))
				logger.Logger.Error(fmt.Sprintf("stacktrace from panic: %s", string(debug.Stack())))
				logger.Logger.Error(fmt.Sprintf("request: %s %s", c.Request.Method, c.Request.URL.Path))
				body, _ := common.GetRequestBody(c)
				logger.Logger.Error(fmt.Sprintf("request body: %s", string(body)))
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": gin.H{
						"message": fmt.Sprintf("Panic detected, error: %v. Please submit an issue with the related log here: https://github.com/Laisky/one-api", err),
						"type":    "one_api_panic",
					},
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
