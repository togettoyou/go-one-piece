package util

import (
	"github.com/gin-gonic/gin"
	"go-one-piece/util/conf"
	"go-one-piece/util/logger"
	"go.uber.org/zap"
)

// 重设配置
func Reset() {
	conf.Reset()
	logger.Reset()
	resetGinMode()
	zap.L().Info("重载配置")
}

const (
	DebugMode   = "debug"
	ReleaseMode = "release"
	TestMode    = "test"
)

func resetGinMode() {
	switch conf.Config.Server.RunMode {
	case DebugMode, "":
		gin.SetMode(DebugMode)
	case ReleaseMode:
		gin.SetMode(ReleaseMode)
	case TestMode:
		gin.SetMode(TestMode)
	default:
		gin.SetMode(DebugMode)
	}
}
