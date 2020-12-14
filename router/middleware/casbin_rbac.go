package middleware

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/service/casbin_service"
	"go-one-server/util/errno"
	"go.uber.org/zap"
)

// 权限验证需要配合jwt使用
func CasbinRBAC() gin.HandlerFunc {
	return func(c *gin.Context) {
		g := Gin{Ctx: c}
		// 获取请求的URI
		path := c.Request.URL.RequestURI()
		// 获取请求方法
		method := c.Request.Method
		// 获取用户名
		claims := GetJWTClaims(c)
		if claims == nil {
			g.SendNoDataResponse(errno.ErrCasbin)
			c.Abort()
			return
		}
		roleKey := casbin_service.GetRoleKeyByUser(claims.Username)
		// 传入casbin请求规则，判断策略中是否存在
		success, err := casbin_service.Casbin().Enforce(roleKey, path, method)
		if err != nil || !success {
			if err != nil {
				zap.L().Error(err.Error())
			}
			g.SendNoDataResponse(errno.ErrPermission)
			c.Abort()
			return
		}
		c.Next()
	}
}
