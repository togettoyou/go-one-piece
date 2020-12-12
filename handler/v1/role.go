package v1

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/model"
)

type roleBody struct {
	RoleName string `json:"role_name" binding:"required,max=20" example:"角色名"`
}

// @Tags 角色
// @Summary 添加角色
// @Produce json
// @Security ApiKeyAuth
// @Param data body roleBody true "角色信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/role [post]
func AddRole(c *gin.Context) {
	g := Gin{Ctx: c}
	var body roleBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	role := model.Role{
		RoleName: body.RoleName,
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
