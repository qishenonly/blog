package global

import (
	"blog/config"
	"gorm.io/gorm"
)

var (
	// Config 是一个全局变量，用来保存配置信息
	Config *config.Config

	// DB 是一个全局变量，用来保存数据库连接
	DB *gorm.DB
)
