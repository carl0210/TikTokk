package middleware

import (
	"TikTokk/internal/pkg/Tlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)

		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		Tlog.Std.Infow(fmt.Sprintf("[GIN] %v |%3d|%13v|%15s|%-7s %#v\n%s",
			startTime.Format("2006/01/02 - 15:04:05.000"),
			status,
			latency,
			clientIP,
			method,
			path,
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		))
	}
}
