package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()
		cost := time.Since(start)
		statusCode := c.Writer.Status()
		data := make([]zap.Field, 0)
		data = append(data,
			zap.Int("statusCode", statusCode),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.RequestURI),
			zap.String("ip", c.ClientIP()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
		if statusCode > 499 {
			zap.L().Error(path, data...)
		} else if statusCode > 399 {
			zap.L().Warn(path, data...)
		} else {
			zap.L().Info(path, data...)
		}
	}
}
