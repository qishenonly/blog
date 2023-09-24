package core

import (
	"blog/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() *gorm.DB {
	// 判断数据库配置信息是否为空
	if global.Config.Mysql.Host == "" && global.Config.Mysql.Port == 0 {
		global.Logger.Warnln("数据库配置信息为空")
		return nil
	}
	dsn := global.Config.Mysql.GetDSN()

	// 设置日志模式
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		// 开发环境下，打印日志
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Silent)
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Logger.Fatalf("[%s] 连接数据库失败: %s", dsn, err)
		return nil
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		global.Logger.Fatalf("设置连接池失败: %s", err)
		return nil
	}
	// 设置空闲连接池中的最大连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour * 3)

	global.Logger.Infof("数据库连接成功")
	return db

}
