package service

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go-one-server/model"
)

const defaultConfigFile = "rbac_model.conf"

func Casbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(model.DB())
	e, _ := casbin.NewEnforcer(defaultConfigFile, a)
	_ = e.LoadPolicy()
	return e
}
