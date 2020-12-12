package service

import (
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go-one-server/model"
	"go-one-server/util/conf"
	"go.uber.org/zap"
)

var (
	adapter  *gormadapter.Adapter
	enforcer *casbin.Enforcer
)

func casbinSetup() {
	var err error
	defer func() {
		panicErr := recover()
		if panicErr != nil {
			zap.S().Fatal(panicErr)
		}
		if err != nil {
			zap.L().Fatal(err.Error())
		}
	}()
	// 将数据库连接同步给Casbin插件， 插件用来操作数据库
	adapter, err = gormadapter.NewAdapterByDB(model.DB())
	if err != nil {
		return
	}
	m, err := casbinmodel.NewModelFromString(conf.Config.Casbin.Model)
	if err != nil {
		return
	}
	enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		return
	}
	// 开启权限认证日志
	enforcer.EnableLog(conf.Config.Casbin.Log)
	// 加载数据库中的策略
	err = enforcer.LoadPolicy()
}

func Casbin() *casbin.Enforcer {
	_ = enforcer.LoadPolicy()
	return enforcer
}
