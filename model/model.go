package model

import (
	"github.com/gin-gonic/gin"
	"go-one-server/util/conf"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func level() logger.LogLevel {
	switch conf.Config.Server.RunMode {
	case gin.ReleaseMode:
		return logger.Error
	default:
		return logger.Info
	}
}

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(conf.Config.Mysql.Dsn), &gorm.Config{
		Logger: logger.Default.LogMode(level()),
	})
	if err != nil {
		zap.L().Error(err.Error())
	}
}

func Reset() {
	db.Config.Logger = logger.Default.LogMode(level())
}
