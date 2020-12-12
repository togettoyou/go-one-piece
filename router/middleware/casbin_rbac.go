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
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 获取用户角色ID
		claims := GetJWTClaims(c)
		if claims == nil {
			g.SendNoDataResponse(errno.ErrTokenFailure)
			c.Abort()
			return
		}
		// 判断策略中是否存在
		success, err := casbin_service.CasbinLoadPolicy().Enforce(claims.RoleID, obj, act)
		if err != nil {
			zap.L().Error(err.Error())
			g.SendNoDataResponse(errno.ErrUnknown)
			c.Abort()
			return
		}
		if !success {
			g.SendNoDataResponse(errno.ErrPermission)
			c.Abort()
			return
		}
		c.Next()
	}
}
