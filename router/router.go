package router

import (
	"github.com/gin-gonic/gin"
	"go-one-piece/router/middleware"
)

//初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	//全局 日志中间件
	r.Use(middleware.LoggerToFile())
	//全局 Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	r.Use(gin.Recovery())
	return r
}
