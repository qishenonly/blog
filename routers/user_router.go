package routers

import "github.com/qishenonly/blog/api"

func (router RouterGroup) InitUserRouter() {
	userApi := api.NewApiGroup().UserApi

	// 首页获取用户信息
	router.GET("/user/info", userApi.GetUserInfo)
}
