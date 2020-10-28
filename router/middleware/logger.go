package middleware

import (
	"bytes"
	"encoding/json"
	"go-one-server/handler"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type respLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w respLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w respLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyLogWriter := &respLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		start := time.Now()
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
		var resp handler.Response
		var result string
		if bodyLogWriter.body.String() != "" {
			err := json.Unmarshal(bodyLogWriter.body.Bytes(), &resp)
			if err == nil {
				result = "\tresponse msg: " + resp.Msg
			}
		}
		if statusCode > 499 {
			zap.L().Error(result+"\n", data...)
		} else if statusCode > 399 {
			zap.L().Warn(result+"\n", data...)
		} else {
			zap.L().Info(result+"\n", data...)
		}
	}
}
