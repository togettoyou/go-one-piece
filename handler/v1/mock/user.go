package mock

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/model"
	"go-one-server/service/mock_service/user_service"
	"go-one-server/util/tools"
)

type RegisteredBody struct {
	Username  string `json:"username" binding:"required,checkUsername"`
	Password  string `json:"password" binding:"required"`
	NickName  string `json:"nick_name"`
	HeaderImg string `json:"header_img" binding:"omitempty,url"`
}

// @Tags mock用户
// @Summary 用户注册
// @Produce json
// @Param data body RegisteredBody true "注册信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/mock/registered [post]
func Registered(c *gin.Context) {
	g := Gin{Ctx: c}
	var body RegisteredBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	salt := tools.RangeString(6)
	user := model.User{
		Username:  body.Username,
		Password:  tools.MD5V(body.Password + salt),
		Salt:      salt,
		NickName:  body.NickName,
		HeaderImg: body.HeaderImg,
	}
	if g.HasError(user.Create()) {
		return
	}
	g.OkWithMsgResponse("注册成功")
}

type LoginBody struct {
	Username string `json:"username" binding:"required,checkUsername"`
	Password string `json:"password" binding:"required"`
}

// @Tags mock用户
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
	data, err := user_service.GetUserInfo(body.Username, body.Password)
	if g.HasError(err) {
		return
	}
	g.OkWithDataResponse(data)
}
