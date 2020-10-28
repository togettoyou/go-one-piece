package model

import (
	"go-one-server/util/conf"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	db, err = gorm.Open(mysql.Open(conf.Config.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error(err.Error())
	}
}
