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

var modes = map[string]string{
	"debug":   "debug",
	"release": "release",
	"test":    "test",
}

func resetGinMode() {
	if mode, ok := modes[conf.Config.Server.RunMode]; ok {
		gin.SetMode(mode)
	} else {
		gin.SetMode("debug")
	}
}
