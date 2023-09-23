package main

import "blog/core"

func main() {
	// 读取配置文件
	core.InitCnf()

	// 初始化日志
	// 初始化数据库
	core.InitGorm()

	// 初始化路由
	// 启动服务
}
