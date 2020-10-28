package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-one-piece/handler/v1/examples"
	"go-one-piece/router/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middleware.Logger())
	// debug模式开启性能分析
	pprof.Register(r)
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
