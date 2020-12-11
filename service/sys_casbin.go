package service

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go-one-server/model"
	"go.uber.org/zap"
)

var DefaultCasbinFile = "rbac_model.conf"

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
	enforcer, err = casbin.NewEnforcer(DefaultCasbinFile, adapter)
	if err != nil {
		return
	}
	// 开启权限认证日志
	enforcer.EnableLog(true)
	// 加载数据库中的策略
	err = enforcer.LoadPolicy()
}

func Casbin() *casbin.Enforcer {
	_ = enforcer.LoadPolicy()
	return enforcer
}
