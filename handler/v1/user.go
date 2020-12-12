package v1

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/model"
	"go-one-server/service/user_service"
	"go-one-server/util/tools"
)

type registeredUserBody struct {
	Username string `json:"username" binding:"required,checkUsername" example:"user1"`
	Password string `json:"password" binding:"required" example:"123456"`
	Remark   string `json:"remark" binding:"omitempty" example:"备注"`
	RoleID   string `json:"role_id" binding:"omitempty" example:"角色ID(可选)"`
}

// @Tags 用户
// @Summary 用户注册
// @Produce json
// @Param data body registeredUserBody true "注册信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/user/registered [post]
func Registered(c *gin.Context) {
	g := Gin{Ctx: c}
	var body registeredUserBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	salt := tools.NewRandom().String(6)
	roleID := sql.NullString{
		Valid: false,
	}
	if body.RoleID != "" {
		roleID = sql.NullString{
			String: body.RoleID,
			Valid:  true,
		}
	}
	user := model.User{
		Username: body.Username,
		Password: tools.MD5V(body.Password + salt),
		Salt:     salt,
		Remark:   body.Remark,
		RoleID:   roleID,
	}
	if g.HasSqlError(user.Create()) {
		return
	}
	g.OkWithMsgResponse("注册成功")
}

type loginBody struct {
	Username string `json:"username" binding:"required,checkUsername" example:"user1"`
	Password string `json:"password" binding:"required" example:"123456"`
}

// @Tags 用户
// @Summary 用户登录
// @Produce json
// @Param data body loginBody true "登录信息"
// @Success 200 {object} handler.Response
// @Failure 500 {object} handler.Response
// @Router /api/v1/user/login [post]
func Login(c *gin.Context) {
	g := Gin{Ctx: c}
	var body loginBody
	if !g.ParseJSONRequest(&body) {
		return
	}
	data, err := user_service.GetUserInfo(body.Username, body.Password)
	if g.HasError(err) {
		return
	}
	g.OkWithDataResponse(data)
}
