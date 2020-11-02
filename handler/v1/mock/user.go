package mock

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/model"
	"go-one-server/router/middleware"
	"go-one-server/service/mock_service/user_service"
	"go-one-server/util/tools"
)

type PaginationQueryBody struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

// @Tags mock用户
// @Summary 分页获取用户列表
// @Produce  json
// @Param page query int false "页码"
// @Param page_size query int false "页面大小"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/mock/userList [get]
func GetUserList(c *gin.Context) {
	g := Gin{Ctx: c}
	var body PaginationQueryBody
	if !g.ParseQueryRequest(&body) {
		return
	}
	userList, err := model.GetUserList(body.Page, body.PageSize)
	if g.HasError(err) {
		return
	}
	g.OkWithDataResponse(userList)
}

type RegisteredBody struct {
	Username  string `json:"username" binding:"required,checkUsername"`
	Password  string `json:"password" binding:"required"`
	NickName  string `json:"nick_name" binding:"omitempty,max=16"`
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
	salt := tools.NewRandom().String(6)
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

// @Tags mock用户
// @Summary 用户查看信息
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/mock/userInfo [get]
func GetUserInfo(c *gin.Context) {
	g := Gin{Ctx: c}
	claims := middleware.GetJWTClaims(c)
	user, err := model.FindUser(claims.Issuer)
	if g.HasError(err) {
		return
	}
	g.OkWithDataResponse(user)
}

type UserInfoBody struct {
	NickName  string `json:"nick_name" binding:"omitempty,max=16"`
	HeaderImg string `json:"header_img" binding:"omitempty,url"`
}

// @Tags mock用户
// @Summary 用户修改信息
// @Produce json
// @Security ApiKeyAuth
// @Param data body UserInfoBody true "修改信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/mock/userInfo [patch]
func UpdateUserInfo(c *gin.Context) {
	g := Gin{Ctx: c}
	var body UserInfoBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	claims := middleware.GetJWTClaims(c)
	user := model.User{
		Username:  claims.Issuer,
		NickName:  body.NickName,
		HeaderImg: body.HeaderImg,
	}
	if g.HasError(user.UpdateUserInfo()) {
		return
	}
	g.OkWithMsgResponse("修改成功")
}
