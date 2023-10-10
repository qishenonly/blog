package api

import (
	"blog/api/auth"
	"blog/api/setting"
)

type ApiGroup struct {
	// SettingApi 是一个 SettingApi 类型的变量，用来调用 setting 包中的方法
	SettingApi setting.SettingApi

	// AuthApi 是一个 AuthApi 类型的变量，用来调用 auth 包中的方法
	AuthApi auth.AuthApi
}

// NewApiGroup 是一个 ApiGroup 类型的函数，用来初始化 ApiGroup 结构体
func NewApiGroup() ApiGroup {
	return ApiGroup{
		SettingApi: setting.SettingApi{},
		AuthApi:    auth.AuthApi{},
	}
}
