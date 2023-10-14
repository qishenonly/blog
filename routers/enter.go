package routers

import (
	"blog/global"
	"blog/middlewares"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	// 设置运行模式
	gin.SetMode(global.Config.System.GetEnv())

	// 初始化路由
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	for _, handlerFunc := range middlewares.GetAllMiddleware() {
		router.Use(handlerFunc)
	}

	// 初始化路由组
	apiRouterGroup := router.Group("/api")
	routerGroup := RouterGroup{apiRouterGroup}

	// 系统配置路由组
	routerGroup.InitSettingRouter()

	// Auth 路由组
	routerGroup.InitAuthRouter()

	// Home 路由组
	routerGroup.InitHomeRouter()

	return router

}
