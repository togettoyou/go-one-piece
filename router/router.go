package router

import (
	"github.com/gin-gonic/gin"
	"go-one-piece/handler/v1/examples"
	"go-one-piece/router/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	apiV1 := r.Group("/api/v1")
	initExamplesRouter(apiV1)
	return r
}

func initExamplesRouter(api *gin.RouterGroup) {
	examplesRouterGroup := api.Group("/examples")
	{
		examplesRouterGroup.GET("/get", examples.GetExamples)
	}
}
