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
// @Router /api/v1/casbin/{roleID} [get]
func GetAllCasbin(c *gin.Context) {
	g := Gin{Ctx: c}
	var body casbinPath
	if !g.ParseUriRequest(&body) {
		return
	}
	pathMaps := casbin_service.GetPolicyPathByRoleID(body.RoleID)
	g.OkWithDataResponse(pathMaps)
}

type casbinBody struct {
	CasbinInReceive model.CasbinInReceive `json:"casbin_in_receive" binding:"required"`
}

// @Tags casbin
// @Summary 更新角色权限
// @Produce  json
// @Security ApiKeyAuth
// @Param data body casbinBody true "权限信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/casbin [put]
func UpdateCasbin(c *gin.Context) {
	g := Gin{Ctx: c}
	var body casbinBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	if g.HasError(casbin_service.UpdateCasbin(body.CasbinInReceive)) {
		return
	}
	pathMaps := casbin_service.GetPolicyPathByRoleID(body.CasbinInReceive.RoleID)
	g.OkWithDataResponse(pathMaps)
}
