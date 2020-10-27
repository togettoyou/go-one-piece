package router

import (
	"github.com/gin-gonic/gin"
	"go-one-piece/router/middleware"
)

//初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	return r
}
