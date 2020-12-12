package casbin_service

import (
	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"go-one-server/model"
	"go-one-server/util/conf"
	"go-one-server/util/errno"
	"go.uber.org/zap"
	"strings"
)

var (
	adapter  *gormadapter.Adapter
	enforcer *casbin.Enforcer
)

// casbin只是映射角色对应的权限，没有校验登录状态的功能，需配合jwt
func Setup() {
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
	// 加载模型
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
	// 添加自定义匹配器
	enforcer.AddFunction("queryMatch", queryMatchFunc)
	// 加载数据库中的策略
	err = enforcer.LoadPolicy()
}

// 如果手动更新数据库，casbin是不会随之更新的，需要重新LoadPolicy
func CasbinLoadPolicy() *casbin.Enforcer {
	_ = enforcer.LoadPolicy()
	return enforcer
}

func Casbin() *casbin.Enforcer {
	return enforcer
}

// 获取用户的角色ID
func GetRoleByUser(username string) (string, error) {
	e := Casbin()
	roles, err := e.GetRolesForUser(username)
	//policy := e.GetFilteredNamedPolicy("g", 0, username)
	//zap.S().Info(policy)
	if err != nil {
		return "", err
	}
	if len(roles) < 1 {
		return "", errno.ErrUserRoleNotFound
	}
	return roles[0], nil
}

// 设置用户角色
func SetUserRole(username, roleID, roleName string) error {
	e := Casbin()
	// 本系统只支持一个用户一个角色，所以先删除再增加
	_, err := e.DeleteRolesForUser(username)
	if err != nil {
		return err
	}
	success, err := e.AddRoleForUser(username, roleID, roleName)
	if err != nil {
		return err
	}
	if success == false {
		return errno.ErrSetCasbinUserRole
	}
	return nil
}

// 获取角色权限列表
func GetApiByRoleID(roleID string) []model.CasbinRoleApiInfo {
	e := Casbin()
	pathMaps := make([]model.CasbinRoleApiInfo, 0)
	list := e.GetFilteredPolicy(0, roleID)
	for _, v := range list {
		pathMaps = append(pathMaps, model.CasbinRoleApiInfo{
			Path:   v[1],
			Method: v[2],
			ApiDes: v[3],
		})
	}
	return pathMaps
}

// 更新角色权限
func UpdateRoleApi(roleID string, casbinInfos []model.CasbinRoleApiInfo) error {
	ClearCasbin(0, roleID)
	var rules [][]string
	for _, v := range casbinInfos {
		rules = append(rules, []string{roleID, v.Path, v.Method, v.ApiDes})
	}
	e := Casbin()
	success, err := e.AddPolicies(rules)
	if err != nil {
		return err
	}
	if success == false {
		return errno.ErrUpdateCasbinRoleApi
	}
	return nil
}

// 清除匹配的权限
func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

func queryMatch(fullNameKey1 string, key2 string) bool {
	//去除路径中?后面的参数
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}

func queryMatchFunc(args ...interface{}) (interface{}, error) {
	key1 := args[0].(string)
	key2 := args[1].(string)
	return queryMatch(key1, key2), nil
}
