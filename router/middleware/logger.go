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
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		statusCode := c.Writer.Status()
		if len(c.Errors) > 0 {
			zap.L().Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode > 499 {
				zap.L().Error(path,
					zap.Int("statusCode", statusCode),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
					zap.Duration("cost", cost),
				)
			} else if statusCode > 399 {
				zap.L().Warn(path,
					zap.Int("statusCode", statusCode),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
					zap.Duration("cost", cost),
				)
			} else {
				zap.L().Info(path,
					zap.Int("statusCode", statusCode),
					zap.String("method", c.Request.Method),
					zap.String("path", path),
					zap.String("query", query),
					zap.String("ip", c.ClientIP()),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
					zap.Duration("cost", cost),
				)
			}
		}
	}
}
