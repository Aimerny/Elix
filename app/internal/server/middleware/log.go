package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	status200 = 42
	status404 = 43
	status500 = 41

	methodGET = 44
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped

		// Stop timer
		end := time.Now()
		timeSub := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		//bodySize := c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}

		var statusColor string
		switch statusCode {
		case 200:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status200, statusCode)
		case 404:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", status404, statusCode)
		}

		var methodColor string
		switch method {
		case "GET":
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodGET, method)

		}

		logrus.Infof("%s |%s| %d | %s | %s | %s",
			start.Format("2006-01-02 15:04:06"),
			statusColor,
			timeSub,
			clientIP,
			methodColor,
			path,
		)

	}
}
