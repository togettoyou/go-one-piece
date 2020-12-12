package main

import (
	"github.com/spf13/pflag"
	"go-one-server/model"
	"go-one-server/service"
	"go-one-server/service/casbin_service"
	"go-one-server/util/conf"
	"go-one-server/util/logger"
	"go-one-server/util/tools"
	"go-one-server/util/validator"
	"go.uber.org/zap"
)

func setup() {
	conf.Setup()
	logger.Setup()
	validator.Setup()
	model.Setup()
	service.Setup()
}

var (
	config = pflag.StringP("config", "c", "config.yaml", "指定配置文件路径")
)

func main() {
	pflag.Parse()
	conf.DefaultConfigFile = *config
	setup()
	initDB()
}

func initDB() {
	initUserDB()
	initRoleDB()
	initUserRole()
	initCasbin()
}

func initUserDB() {
	users := []model.User{
		{
			Username: "admin",
			Password: tools.MD5V("123456" + "ABCDEF"),
			Salt:     "ABCDEF",
			Remark:   "超级管理员",
		}, {
			Username: "kefu1",
			Password: tools.MD5V("123456" + "ABCDEF"),
			Salt:     "ABCDEF",
			Remark:   "客服1号",
		}, {
			Username: "user1",
			Password: tools.MD5V("123456" + "ABCDEF"),
			Salt:     "ABCDEF",
			Remark:   "普通用户1",
		},
	}
	err := model.DB().Create(users).Error
	if err != nil {
		zap.L().Error(err.Error())
	}
}

func initRoleDB() {
	roles := []model.Role{
		{
			RoleInfo: model.RoleInfo{
				RoleID:   "root",
				RoleName: "超级管理员",
			},
		}, {
			RoleInfo: model.RoleInfo{
				RoleID:   "kefu",
				RoleName: "客服",
			},
		}, {
			RoleInfo: model.RoleInfo{
				RoleID:   "member",
				RoleName: "会员",
			},
		},
	}
	err := model.DB().Create(roles).Error
	if err != nil {
		zap.L().Error(err.Error())
	}
}

func initUserRole() {
	err := casbin_service.SetUserRole("admin", "root", "超级管理员")
	if err != nil {
		zap.L().Error(err.Error())
	}
	err = casbin_service.SetUserRole("kefu1", "kefu", "客服")
	if err != nil {
		zap.L().Error(err.Error())
	}
	err = casbin_service.SetUserRole("user1", "member", "会员")
	if err != nil {
		zap.L().Error(err.Error())
	}
}

func initCasbin() {
	err := casbin_service.UpdateRoleApi("root", []model.CasbinRoleApiInfo{{
		Path:   "/*",
		Method: "*",
		ApiDes: "所有接口",
	}})
	if err != nil {
		zap.L().Error(err.Error())
	}
	err = casbin_service.UpdateRoleApi("kefu", []model.CasbinRoleApiInfo{{
		Path:   "/api/v1/kefu/*",
		Method: "*",
		ApiDes: "kefu下所有接口",
	}})
	if err != nil {
		zap.L().Error(err.Error())
	}
	err = casbin_service.UpdateRoleApi("member", []model.CasbinRoleApiInfo{{
		Path:   "/api/v1/member/*",
		Method: "*",
		ApiDes: "member下所有接口",
	}})
	if err != nil {
		zap.L().Error(err.Error())
	}
}
