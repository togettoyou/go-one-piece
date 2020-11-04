package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go-one-server/handler/v1/mock"
	"go-one-server/router/middleware"
)

var swagHandler gin.HandlerFunc

func HasDocs() bool {
	return swagHandler != nil
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(middleware.Logger())
	//debug模式开启性能分析
	pprof.Register(r)
	//swagger文档，根据build tag控制编译减少二进制文件大小
	if HasDocs() {
		r.GET("/swagger/*any", swagHandler)
	}
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
		mockRouterGroup.GET("/userList", mock.GetUserList)
		mockRouterGroup.POST("/registered", mock.Registered)
		mockRouterGroup.POST("/login", mock.Login)
		mockRouterGroup.GET("/userInfo", middleware.JWT(), mock.GetUserInfo)
		mockRouterGroup.PATCH("/userInfo", middleware.JWT(), mock.UpdateUserInfo)
	}
}
