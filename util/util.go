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
	zap.S().Info("config change", conf.Config)
}

const (
	DebugMode   string = "debug"
	ReleaseMode string = "release"
	TestMode    string = "test"
)

var modes = map[string]string{
	"debug":   DebugMode,
	"release": ReleaseMode,
	"test":    TestMode,
}

func resetGinMode() {
	if mode, ok := modes[conf.Config.Server.RunMode]; ok {
		gin.SetMode(mode)
	} else {
		gin.SetMode(DebugMode)
	}
}
