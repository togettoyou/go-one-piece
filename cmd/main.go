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
	// api权限
	apis := []*model.Api{
		{ApiInfo: model.ApiInfo{Path: "/api/v1/role", Method: "POST", ApiGroup: "role", Description: "添加角色"}},
		{ApiInfo: model.ApiInfo{Path: "/api/v1/role", Method: "GET", ApiGroup: "role", Description: "分页获取角色列表"}},
		{ApiInfo: model.ApiInfo{Path: "/api/v1/role/:role_key", Method: "DELETE", ApiGroup: "role", Description: "删除角色"}},

		{ApiInfo: model.ApiInfo{Path: "/api/v1/api", Method: "POST", ApiGroup: "api", Description: "添加api"}},
		{ApiInfo: model.ApiInfo{Path: "/api/v1/api", Method: "GET", ApiGroup: "api", Description: "分页获取api"}},
		{ApiInfo: model.ApiInfo{Path: "/api/v1/api/:id", Method: "DELETE", ApiGroup: "api", Description: "删除api"}},

		{ApiInfo: model.ApiInfo{Path: "/api/v1/setUserRole", Method: "POST", ApiGroup: "casbin", Description: "设置用户角色"}},
		{ApiInfo: model.ApiInfo{Path: "/api/v1/casbin/api/:role_key", Method: "GET", ApiGroup: "casbin", Description: "查看角色权限"}},
		{ApiInfo: model.ApiInfo{Path: "/api/v1/casbin/api/:role_key", Method: "PUT", ApiGroup: "casbin", Description: "更新角色权限"}},
	}
	model.DB().Create(apis)

	// 角色
	roles := []model.Role{
		{RoleInfo: model.RoleInfo{RoleKey: "root", RoleName: "系统管理员", Remark: "系统管理员拥有所有权限"}},
		{RoleInfo: model.RoleInfo{RoleKey: "member", RoleName: "会员", Remark: "会员只有部分权限"}},
	}
	model.DB().Create(roles)

	// 角色添加api权限
	casbin_service.UpdateRoleApi("member", []model.Api{*apis[1], *apis[4], *apis[7]})

	// 用户
	users := []model.User{
		{Username: "admin", Password: tools.MD5V("123456" + "ABCDEF"), Salt: "ABCDEF", Remark: "管理员"},
		{Username: "user1", Password: tools.MD5V("123456" + "ABCDEF"), Salt: "ABCDEF", Remark: "用户1"},
	}
	model.DB().Create(users)

	// 设置用户角色
	casbin_service.SetUserRole("admin", "root")
	casbin_service.SetUserRole("user1", "member")
}
