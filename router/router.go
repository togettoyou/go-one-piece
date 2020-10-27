package router

import (
	"github.com/gin-gonic/gin"
	"go-one-piece/router/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	return r
}
