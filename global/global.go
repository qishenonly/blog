package global

import (
	"github.com/qishenonly/blog/config"

	"github.com/ByteStorage/FlyDB/engine"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	// Config 是一个全局变量，用来保存配置信息
	Config *config.Config

	// DB 是一个全局变量，用来保存数据库连接
	DB *gorm.DB

	// Logger 是一个全局变量，用来保存日志记录器
	Logger *logrus.Logger

	// Cache 是一个全局变量，用来保存缓存
	Cache *engine.DB
)
