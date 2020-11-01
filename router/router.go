package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-one-server/docs"
	"go-one-server/handler/v1/mock"
	"go-one-server/router/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middleware.Logger())
	//debug模式开启性能分析
	pprof.Register(r)
	//swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//api路由分组v1版本
	apiV1 := r.Group("/api/v1")
	initMockRouter(apiV1)
	return r
}

func initMockRouter(api *gin.RouterGroup) {
	mockRouterGroup := api.Group("/mock")
	{
		mockRouterGroup.GET("/get", mock.Get)
		mockRouterGroup.GET("/uri/:id", mock.Uri)
		mockRouterGroup.GET("/query", mock.Query)
		mockRouterGroup.POST("/form", mock.FormData)
		mockRouterGroup.POST("/json", mock.JSON)
		mockRouterGroup.GET("/query/array", mock.QueryArray)
		mockRouterGroup.GET("/query/map", mock.QueryMap)
	}
	{
		mockRouterGroup.POST("/registered", mock.Registered)
		mockRouterGroup.POST("/login", mock.Login)
		mockRouterGroup.PATCH("/userInfo", middleware.JWT(), mock.UpdateUserInfo)
	}
}
