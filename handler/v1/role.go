package v1

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/model"
	"go-one-server/service/casbin_service"
	"go-one-server/util/errno"
)

// @Tags 角色
// @Summary 添加角色
// @Produce json
// @Security ApiKeyAuth
// @Param data body model.RoleInfo true "角色信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/role [post]
func AddRole(c *gin.Context) {
	g := Gin{Ctx: c}
	var body model.RoleInfo
	if !g.ParseJSONRequest(&body) {
		return
	}
	role := model.Role{
		RoleInfo: body,
	}
	if g.HasSqlError(role.Create()) {
		return
	}
	g.OkWithMsgResponse("添加成功")
}

type PaginationQueryBody struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

// @Tags 角色
// @Summary 分页获取角色列表
// @Produce  json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param page_size query int false "页面大小"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/role [get]
func GetRoleList(c *gin.Context) {
	g := Gin{Ctx: c}
	var body PaginationQueryBody
	if !g.ParseQueryRequest(&body) {
		return
	}
	roleList, err := model.GetRoleList(body.Page, body.PageSize)
	if g.HasSqlError(err) {
		return
	}
	g.OkWithDataResponse(roleList)
}

type RolePath struct {
	RoleKey string `json:"role_key" uri:"role_key" binding:"required"`
}

// @Tags 角色
// @Summary 删除角色
// @Produce  json
// @Security ApiKeyAuth
// @Param role_key path string true "角色代码"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/role/{role_key} [delete]
func DelRole(c *gin.Context) {
	g := Gin{Ctx: c}
	var uri RolePath
	if !g.ParseUriRequest(&uri) {
		return
	}
	// 清空角色拥有的权限
	if casbin_service.ClearRoleApi(uri.RoleKey) {
		if g.HasSqlError(model.DelRole(uri.RoleKey)) {
			return
		}
		g.OkWithMsgResponse("删除成功，权限已清空")
		return
	} else {
		g.SendNoDataResponse(errno.ErrDelRoleApi)
	}
}
