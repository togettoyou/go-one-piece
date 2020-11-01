package model

import "go.uber.org/zap"

func initTable(tables ...interface{}) {
	for _, table := range tables {
		// 如果数据库中不存在该表，则创建表
		if !db.Migrator().HasTable(table) {
			// 创建表时添加后缀
			if err := db.Set("gorm:table_options",
				"ENGINE=InnoDB DEFAULT CHARSET=utf8").
				Migrator().CreateTable(table); err != nil {
				zap.L().Error(err.Error())
			}
		}
	}
}
