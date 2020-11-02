package model

import (
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

var logMode = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

func level() logger.LogLevel {
	if logLevel, ok := logMode[conf.Config.Mysql.LogMode]; ok {
		return logLevel
	} else {
		return logger.Silent
	}
}

func Setup() {
	var err error
	db, err = gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       conf.Config.Mysql.Dsn,
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}),
		&gorm.Config{
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
