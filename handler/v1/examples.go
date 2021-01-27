package v1

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
)

// @Tags examples
// @Summary Get请求
// @Produce  json
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/get [get]
func Get(c *gin.Context) {
	g := Gin{Ctx: c}
	g.OkResponse()
}

type UriBody struct {
	ID uint `json:"id" uri:"id" binding:"required,min=10"`
}

// @Tags examples
// @Summary uri参数请求
// @Description 路径参数，匹配 /uri/{id}
// @Produce  json
// @Param id path int false "id值"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/uri/{id} [get]
func Uri(c *gin.Context) {
	g := Gin{Ctx: c}
	//id := c.Param("id")
	var body UriBody
	if !g.ParseUriRequest(&body) {
		return
	}
	g.OkWithDataResponse(body)
}

type QueryBody struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

// @Tags examples
// @Summary query参数查询
// @Description 查询参数，匹配 query?id=xxx
// @Produce  json
// @Param email query string true "邮箱"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/query [get]
func Query(c *gin.Context) {
	g := Gin{Ctx: c}
	//email := c.Query("email")
	var body QueryBody
	if !g.ParseQueryRequest(&body) {
		return
	}
	g.OkWithDataResponse(body)
}

type FormBody struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}

// @Tags examples
// @Summary form表单请求
// @Description 处理application/x-www-form-urlencoded类型的POST请求
// @Accept application/x-www-form-urlencoded
// @Produce  json
// @Param email formData string true "邮箱"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/form [post]
func FormData(c *gin.Context) {
	g := Gin{Ctx: c}
	//email := c.PostForm("email")
	var body FormBody
	if !g.ParseFormRequest(&body) {
		return
	}
	g.OkWithDataResponse(body)
}

type JSONBody struct {
	Email    string `json:"email" binding:"required,email" example:"admin@qq.com"`
	Username string `json:"username" binding:"required,checkUsername" example:"admin"`
}

// @Tags examples
// @Summary JSON参数请求
// @Description 邮箱、用户名校验
// @Produce  json
// @Param data body JSONBody true "测试请求json参数"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/json [post]
func JSON(c *gin.Context) {
	g := Gin{Ctx: c}
	var body JSONBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	g.OkWithDataResponse(body)
}

// @Tags examples
// @Summary queryArray参数查询
// @Description 数组参数，匹配多选业务如 array?ids=xxx&ids=xxx&ids=xxx,key一样，value不同
// @Produce  json
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/query/array [get]
func QueryArray(c *gin.Context) {
	g := Gin{Ctx: c}
	ids := c.QueryArray("ids")
	g.OkWithDataResponse(ids)
}

// @Tags examples
// @Summary queryMap参数查询
// @Description map参数，字典参数，匹配 map?ids[a]=123&ids[b]=456&ids[c]=789
// @Produce  json
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/query/map [get]
func QueryMap(c *gin.Context) {
	g := Gin{Ctx: c}
	ids := c.QueryMap("ids")
	g.OkWithDataResponse(ids)
}
