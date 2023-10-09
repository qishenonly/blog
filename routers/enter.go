package routers

import (
	"blog/global"
	"blog/middlewares"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	// 设置运行模式
	gin.SetMode(global.Config.System.GetEnv())

	// 初始化路由
	router := gin.Default()
	for _, handlerFunc := range middlewares.GetAllMiddleware() {
		router.Use(handlerFunc)
	}

	// 初始化路由组
	apiRouterGroup := router.Group("/api")
	routerGroup := RouterGroup{apiRouterGroup}

	// 系统配置路由组
	routerGroup.InitSettingRouter()

	return router

}
