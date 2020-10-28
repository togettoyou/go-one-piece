package model

import (
	"github.com/gin-gonic/gin"
	"go-one-server/util/conf"
	"go-one-server/util/times"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt times.FormatTime
	UpdatedAt times.FormatTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
}

func Reset() {
	db.Config.Logger = logger.Default.LogMode(level())
}
