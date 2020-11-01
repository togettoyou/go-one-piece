package mock

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
)

type LoginBody struct {
	Username string `json:"username" binding:"required,checkUsername"`
	Password string `json:"password" binding:"required"`
}

// @Tags mock
// @Summary 用户登录
// @Produce json
// @Param data body LoginBody true "登录信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/mock/login [post]
func Login(c *gin.Context) {
	g := Gin{Ctx: c}
	var body LoginBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	g.OkResponse()
}
