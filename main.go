package main

import (
	"blog/core"
	"blog/global"
)

func main() {
	// 读取配置文件
	core.InitCnf()

	// 初始化日志
	global.Logger = core.InitLogger()
	global.Logger.Infof("日志记录器初始化成功")

	// 初始化数据库
	global.DB = core.InitGorm()

	// 初始化路由
	// 启动服务
}
