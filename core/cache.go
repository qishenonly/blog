package core

import (
	"github.com/qishenonly/blog/global"

	"github.com/ByteStorage/FlyDB/config"
	"github.com/ByteStorage/FlyDB/engine"
)

func InitCache() *engine.DB {
	if global.Config.Cache.DirPath == "" {
		global.Logger.Warnln("缓存配置信息为空")
		return nil
	}

	option := config.DefaultOptions
	option.DirPath = global.Config.Cache.DirPath
	option.DataFileSize = global.Config.Cache.DataFileSize
	option.SyncWrite = global.Config.Cache.SyncWrite
	option.FIOType = 1

	db, err := engine.NewDB(option)
	if err != nil {
		global.Logger.Errorf("缓存初始化失败：%v", err)
		return nil
	}

	global.Logger.Infof("缓存初始化成功")
	return db
}
