package v1

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/model"
	"go-one-server/service/casbin_service"
)

// @Tags casbin
// @Summary 查看角色权限
// @Produce  json
// @Security ApiKeyAuth
// @Param role_key path string true "角色代码"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/casbin/api/{role_key} [get]
func GetCasbinApi(c *gin.Context) {
	g := Gin{Ctx: c}
	var uri RolePath
	if !g.ParseUriRequest(&uri) {
		return
	}
	apiMaps := casbin_service.GetApiByRoleKey(uri.RoleKey)
	g.OkWithDataResponse(apiMaps)
}

type casbinApiBody struct {
	ApiIDList []uint `json:"api_id_list" binding:"required"`
}

// @Tags casbin
// @Summary 更新角色权限
// @Produce  json
// @Security ApiKeyAuth
// @Param role_key path string true "角色代码"
// @Param data body casbinApiBody true "权限信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/casbin/api/{role_key} [put]
func UpdateCasbinApi(c *gin.Context) {
	g := Gin{Ctx: c}
	var uri RolePath
	if !g.ParseUriRequest(&uri) {
		return
	}
	var body casbinApiBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	// 先从数据库查询角色，确保角色存在
	role, err := model.FindRoleByKey(uri.RoleKey)
	if g.HasSqlError(err) {
		return
	}
	// 从数据库查找权限
	apis, err := model.FindApiInID(body.ApiIDList)
	if g.HasSqlError(err) {
		return
	}
	// 写入casbin
	if g.HasError(casbin_service.UpdateRoleApi(role.RoleKey, apis)) {
		return
	}
	apiMaps := casbin_service.GetApiByRoleKey(uri.RoleKey)
	g.OkWithDataResponse(apiMaps)
}

type userRoleBody struct {
	Username string `json:"username" binding:"required" example:"user1"`
	RoleKey  string `json:"role_key" binding:"required"`
}

// @Tags casbin
// @Summary 设置用户角色
// @Produce json
// @Security ApiKeyAuth
// @Param data body userRoleBody true "用户角色信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/casbin/setUserRole [post]
func SetUserRole(c *gin.Context) {
	g := Gin{Ctx: c}
	var body userRoleBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	// 先从数据库查询角色
	role, err := model.FindRoleByKey(body.RoleKey)
	if g.HasSqlError(err) {
		return
	}
	// 从数据库查询用户
	user, err := model.FindUser(body.Username)
	if g.HasSqlError(err) {
		return
	}
	if g.HasError(casbin_service.SetUserRole(user.Username, role.RoleKey)) {
		return
	}
	g.OkWithMsgResponse("修改成功")
}
