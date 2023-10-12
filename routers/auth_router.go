package routers

import (
	"blog/api"
)

func (router RouterGroup) InitAuthRouter() {
	authApi := api.NewApiGroup().AuthApi
	// 注册
	router.POST("/auth/register", authApi.Register)

	// 注册验证码
	router.GET("/auth/register/code", authApi.RegisterCode)

	// 激活账号
	router.GET("/auth/verify/:token/active/:email", authApi.Verify)

	// 登录
	router.POST("/auth/login", authApi.Login)

	// 登录验证码
	router.GET("/auth/login/code", authApi.LoginCode)

	// 第三方登录
	router.GET("/auth/github", authApi.AuthGithub)
	router.GET("/auth/github_callback", authApi.AuthGithubCallback)
}
