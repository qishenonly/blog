package routers

import (
	"blog/global"
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

	// 初始化路由组
	apiRouterGroup := router.Group("/api")
	routerGroup := RouterGroup{apiRouterGroup}

	// 系统配置路由组
	routerGroup.InitSettingRouter()

	return router

}
