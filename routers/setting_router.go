package routers

import (
	"blog/api"
)

// InitSettingRouter 是一个 RouterGroup 类型的方法，用来初始化路由组
func (router RouterGroup) InitSettingRouter() {
	settingApi := api.NewApiGroup().SettingApi
	router.GET("/setting", settingApi.SettingInfoView)
}
