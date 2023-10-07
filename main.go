package main

import (
	"blog/cmd"
	"blog/core"
	"blog/global"
	"blog/routers"
)

func main() {
	// 读取配置文件
	core.InitCnf()

	// 初始化日志
	global.Logger = core.InitLogger()
	global.Logger.Infof("日志记录器初始化成功")

	// 初始化数据库
	global.DB = core.InitGorm()

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
