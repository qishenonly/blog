package main

import (
	"github.com/qishenonly/blog/cmd"
	"github.com/qishenonly/blog/core"
	_ "github.com/qishenonly/blog/docs"
	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/routers"
)

// @title Gin-Blog API 文档
// @description Gin-Blog API 文档
// @version 0.1
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// 读取配置文件
	core.InitCnf()

	// 初始化日志
	global.Logger = core.InitLogger()
	global.Logger.Infof("日志记录器初始化成功")

	// 初始化数据库
	global.DB = core.InitGorm()

	// 初始化缓存
	global.Cache = core.InitCache()

	//
	option := cmd.ParseOptions()
	if cmd.IsStopWebServer(*option) {
		cmd.InitOptions(*option)
		return
	}

	// 初始化路由
	router := routers.InitRouter()

	// 启动服务
	global.Logger.Infof("服务启动成功，监听端口：%s", global.Config.System.GetAddr())
	err := router.Run(global.Config.System.GetAddr())
	if err != nil {
		global.Logger.Errorf("服务启动失败：%v", err)
	}
}
