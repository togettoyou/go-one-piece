package model

import (
	"github.com/gin-gonic/gin"
	"go-one-server/util/conf"
	"go-one-server/util/tools"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type Model struct {
	ID        uint             `json:"id" gorm:"primarykey"`
	CreatedAt tools.FormatTime `json:"created_at"`
	UpdatedAt tools.FormatTime `json:"-"`
	DeletedAt gorm.DeletedAt   `json:"-" gorm:"index"`
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
			TablePrefix:   "sys_", // 表名前缀，`User` 的表名应该是 `sys_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `sys_user`
		},
	})
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	setConnectionPool()
	initTable(&User{})
}

// 设置连接池
func setConnectionPool() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			zap.L().Error(err.Error())
			return
		}
		maxIdle := conf.Config.Mysql.MaxIdle
		maxOpen := conf.Config.Mysql.MaxOpen
		maxLifetime := conf.Config.Mysql.MaxLifetime
		if maxIdle < 1 {
			maxIdle = 10
		}
		if maxOpen < 1 {
			maxOpen = 100
		}
		if maxLifetime < time.Second {
			maxLifetime = 60 * time.Minute
		}
		sqlDB.SetMaxIdleConns(maxIdle)
		sqlDB.SetMaxOpenConns(maxOpen)
		sqlDB.SetConnMaxLifetime(maxLifetime)
	}
}

func Reset() {
	db.Config.Logger = logger.Default.LogMode(level())
	setConnectionPool()
}
