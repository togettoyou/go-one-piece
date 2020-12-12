package v1

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/model"
	"go-one-server/service/casbin_service"
)

type casbinPath struct {
	RoleID string `json:"roleID" uri:"roleID" binding:"required"`
}

// @Tags casbin
// @Summary 根据角色获取权限列表
// @Produce  json
// @Security ApiKeyAuth
// @Param roleID path string true "角色ID"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/casbin/api/{roleID} [get]
func GetAllCasbinApi(c *gin.Context) {
	g := Gin{Ctx: c}
	var uri casbinPath
	if !g.ParseUriRequest(&uri) {
		return
	}
	pathMaps := casbin_service.GetApiByRoleID(uri.RoleID)
	g.OkWithDataResponse(pathMaps)
}

type casbinBody struct {
	CasbinRoleApiInfo []model.CasbinRoleApiInfo `json:"casbin_role_api_info" binding:"required"`
}

// @Tags casbin
// @Summary 更新角色权限
// @Produce  json
// @Security ApiKeyAuth
// @Param roleID path string true "角色ID"
// @Param data body casbinBody true "权限信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/casbin/api/{roleID} [put]
func UpdateCasbinApi(c *gin.Context) {
	g := Gin{Ctx: c}
	var uri casbinPath
	if !g.ParseUriRequest(&uri) {
		return
	}
	var body casbinBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	// 先从数据库查询角色，确保角色存在
	role, err := model.FindRole(uri.RoleID)
	if g.HasSqlError(err) {
		return
	}
	if g.HasError(casbin_service.UpdateRoleApi(role.RoleID, body.CasbinRoleApiInfo)) {
		return
	}
	pathMaps := casbin_service.GetApiByRoleID(uri.RoleID)
	g.OkWithDataResponse(pathMaps)
}

type userRoleBody struct {
	Username string `json:"username" binding:"required" example:"user1"`
	RoleID   string `json:"roleID" binding:"required" example:"roleID"`
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
	role, err := model.FindRole(body.RoleID)
	if g.HasSqlError(err) {
		return
	}
	if g.HasError(casbin_service.SetUserRole(body.Username, role.RoleID, role.RoleName)) {
		return
	}
	g.OkWithMsgResponse("修改成功")
}
