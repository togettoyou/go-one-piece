package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	v1 "go-one-server/handler/v1"
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
	//开启性能分析
	//实际可以根据需要使用pprof.RouteRegister()控制访问权限
	pprof.Register(r)
	//swagger文档，根据build tag控制编译减少二进制文件大小
	if HasDocs() {
		r.GET("/swagger/*any", swagHandler)
	}
	//api路由分组v1版本
	apiV1 := r.Group("/api/v1")
	initExamplesRouter(apiV1)
	initUserRouter(apiV1)
	initRoleRouter(apiV1)
	initApiRouter(apiV1)
	initCasbinRouter(apiV1)
	return r
}

func initExamplesRouter(api *gin.RouterGroup) {
	examplesRouterGroup := api.Group("/examples")
	{
		examplesRouterGroup.GET("/get", v1.Get)
		examplesRouterGroup.GET("/uri/:id", v1.Uri)
		examplesRouterGroup.GET("/query", v1.Query)
		examplesRouterGroup.POST("/form", v1.FormData)
		examplesRouterGroup.POST("/json", v1.JSON)
		examplesRouterGroup.GET("/query/array", v1.QueryArray)
		examplesRouterGroup.GET("/query/map", v1.QueryMap)
	}
}

func initUserRouter(api *gin.RouterGroup) {
	userRouterGroup := api.Group("/user")
	{
		userRouterGroup.POST("/registered", v1.Registered)
		userRouterGroup.POST("/login", v1.Login)
	}
}

func initRoleRouter(api *gin.RouterGroup) {
	roleRouterGroup := api.Group("/role")
	roleRouterGroup.Use(middleware.JWT()).Use(middleware.CasbinRBAC())
	{
		roleRouterGroup.POST("", v1.AddRole)
		roleRouterGroup.GET("", v1.GetRoleList)
		roleRouterGroup.DELETE("/:role_key", v1.DelRole)
	}
}

func initApiRouter(api *gin.RouterGroup) {
	apiRouterGroup := api.Group("/api")
	apiRouterGroup.Use(middleware.JWT()).Use(middleware.CasbinRBAC())
	{
		apiRouterGroup.GET("", v1.GetApiList)
		apiRouterGroup.POST("", v1.AddApi)
		apiRouterGroup.DELETE("/:id", v1.DelApi)
	}
}

func initCasbinRouter(api *gin.RouterGroup) {
	casbinRouterGroup := api.Group("/casbin")
	casbinRouterGroup.Use(middleware.JWT()).Use(middleware.CasbinRBAC())
	{
		casbinRouterGroup.POST("/setUserRole", v1.SetUserRole)
		casbinRouterGroup.GET("/api/:role_key", v1.GetCasbinApi)
		casbinRouterGroup.PUT("/api/:role_key", v1.UpdateCasbinApi)
	}
}
