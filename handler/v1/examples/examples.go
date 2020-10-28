package examples

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
func GetExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	g.OkResponse()
}

type UriBodyExamples struct {
	ID uint `json:"id" uri:"id" validate:"required,min=10"`
}

// @Tags examples
// @Summary uri参数请求
// @Description 路径参数，匹配 /uri/{id}
// @Produce  json
// @Param id path int false "id值"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/uri/{id} [get]
func UriExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	//id := c.Param("id")
	var body UriBodyExamples
	if !g.ParseUriRequest(&body) {
		return
	}
	g.OkWithDataResponse(body)
}

type QueryBodyExamples struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

// @Tags examples
// @Summary query参数查询
// @Description 查询参数，匹配 query?id=xxx
// @Produce  json
// @Param email query string true "邮箱"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/query [get]
func QueryExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	//email := c.Query("email")
	var body QueryBodyExamples
	if !g.ParseQueryRequest(&body) {
		return
	}
	g.OkWithDataResponse(body)
}

type FormBodyExamples struct {
	Email string `json:"email" form:"email" validate:"required,email"`
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
func FormDataExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	//email := c.PostForm("email")
	var body FormBodyExamples
	if !g.ParseFormRequest(&body) {
		return
	}
	g.OkWithDataResponse(body)
}

type JSONBodyExamples struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,checkUsername"`
}

// @Tags examples
// @Summary JSON参数请求
// @Description 邮箱、用户名校验
// @Produce  json
// @Param data body JSONBodyExamples true "测试请求json参数"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/examples/json [post]
func JSONExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	var body JSONBodyExamples
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
func QueryArrayExamples(c *gin.Context) {
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
func QueryMapExamples(c *gin.Context) {
	g := Gin{Ctx: c}
	ids := c.QueryMap("ids")
	g.OkWithDataResponse(ids)
}
