package middleware

import (
	"github.com/gin-gonic/gin"
	. "go-one-server/handler"
	"go-one-server/service"
	"go-one-server/util/errno"
	"go.uber.org/zap"
)

func CasbinRBAC() gin.HandlerFunc {
	return func(c *gin.Context) {
		g := Gin{Ctx: c}
		// 获取请求的URI
		obj := c.Request.URL.RequestURI()
		// 获取请求方法
		act := c.Request.Method
		// 判断策略中是否存在
		ok, err := service.Casbin().Enforce("root", obj, act)
		if err != nil {
			zap.L().Error(err.Error())
			g.SendNoDataResponse(errno.ErrUnknown)
			c.Abort()
			return
		}
		if !ok {
			g.SendNoDataResponse(errno.ErrPermission)
			c.Abort()
			return
		}
	}
}
