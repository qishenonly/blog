package api

import (
	"blog/api/article"
	"blog/api/auth"
	"blog/api/home"
	"blog/api/setting"
)

type ApiGroup struct {
	// SettingApi 是一个 SettingApi 类型的变量，用来调用 setting 包中的方法
	SettingApi setting.SettingApi

	// AuthApi 是一个 AuthApi 类型的变量，用来调用 auth 包中的方法
	AuthApi auth.AuthApi

	// HomeApi 是一个 HomeApi 类型的变量，用来调用 home 包中的方法
	HomeApi home.HomeApi

	// ArticleApi 是一个 ArticleApi 类型的变量，用来调用 article 包中的方法
	ArticleApi article.ArticleApi
}

// NewApiGroup 是一个 ApiGroup 类型的函数，用来初始化 ApiGroup 结构体
func NewApiGroup() ApiGroup {
	return ApiGroup{
		SettingApi: setting.SettingApi{},
		AuthApi:    auth.AuthApi{},
		HomeApi:    home.HomeApi{},
		ArticleApi: article.ArticleApi{},
	}
}
