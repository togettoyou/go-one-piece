package v1

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/model"
	"go-one-server/service/casbin_service"
	"strconv"
)

// @Tags API
// @Summary 添加API
// @Produce json
// @Security ApiKeyAuth
// @Param data body model.ApiInfo true "api信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/api [post]
func AddApi(c *gin.Context) {
	g := Gin{Ctx: c}
	var body model.ApiInfo
	if !g.ParseJSONRequest(&body) {
		return
	}
	api := model.Api{
		ApiInfo: body,
	}
	if g.HasSqlError(api.Create()) {
		return
	}
	g.OkWithMsgResponse("添加成功")
}

// @Tags API
// @Summary 分页获取API列表
// @Produce  json
// @Security ApiKeyAuth
// @Param page query int false "页码"
// @Param page_size query int false "页面大小"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/api [get]
func GetApiList(c *gin.Context) {
	g := Gin{Ctx: c}
	var body PaginationQueryBody
	if !g.ParseQueryRequest(&body) {
		return
	}
	apiList, err := model.GetApiList(body.Page, body.PageSize)
	if g.HasSqlError(err) {
		return
	}
	g.OkWithDataResponse(apiList)
}

type ApiPath struct {
	ID uint `json:"id" uri:"id" binding:"required"`
}

// @Tags API
// @Summary 删除API
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "api ID"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/api/{id} [delete]
func DelApi(c *gin.Context) {
	g := Gin{Ctx: c}
	var uri ApiPath
	if !g.ParseUriRequest(&uri) {
		return
	}
	// 从数据库查找权限
	api, err := model.FindApiByID(uri.ID)
	if g.HasSqlError(err) {
		return
	}
	// 删除
	if g.HasSqlError(model.DelApi(uri.ID)) {
		return
	}
	// 清空角色拥有的权限
	casbin_service.ClearByApiID(strconv.Itoa(int(api.ID)))
	g.OkWithMsgResponse("删除API成功，相应权限已清空")
}
