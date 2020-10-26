package util

import (
	"github.com/gin-gonic/gin"
	"go-one-piece/util/conf"
	"go-one-piece/util/logging"
)

func Reset() {
	conf.Reset()
	logging.Reset()
	resetGinMode()
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
