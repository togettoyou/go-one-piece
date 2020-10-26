package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-one-piece/util/conf"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	fileName := conf.Config.Logger.LoggerFile
	if fileName != "" {
		logWriter, err := rotatelogs.New(
			// 分割后的文件名称
			fileName+".%Y%m%d.log",
			// 生成软链，指向最新日志文件
			rotatelogs.WithLinkName(fileName),
			// 设置最大保存时间
			rotatelogs.WithMaxAge(conf.Config.Logger.LoggerFileMaxAge*time.Hour),
			// 设置日志切割时间间隔
			rotatelogs.WithRotationTime(conf.Config.Logger.LoggerFileRotationTime*time.Hour),
		)
		if err == nil {
			writeMap := lfshook.WriterMap{
				logrus.DebugLevel: logWriter,
				logrus.InfoLevel:  logWriter,
				logrus.WarnLevel:  logWriter,
				logrus.ErrorLevel: logWriter,
				logrus.FatalLevel: logWriter,
				logrus.PanicLevel: logWriter,
			}
			logrus.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			}))
		}
	}
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 状态码
		statusCode := c.Writer.Status()
		// 日志格式
		entry := logrus.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"uri":        c.Request.RequestURI,
			"ip":         c.ClientIP(),
			"statusCode": statusCode,
			"costTime":   endTime.Sub(startTime),
		})
		// 返回数据
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode > 499 {
				entry.Error()
			} else if statusCode > 399 {
				entry.Warn()
			} else {
				entry.Info()
			}
		}
	}
}
