package routers

import (
	"github.com/qishenonly/blog/api"
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

	// 重置密码
	router.POST("/auth/reset_pwd/:email", authApi.ResetPwd)

	// 重置密码验证码
	router.GET("/auth/reset_pwd/code", authApi.ResetPwdCode)

	// 重置密码激活
	router.GET("/auth/verify/reset_pwd/:token/active/:email", authApi.VerifyResetPwd)

	// 退出登录
	router.POST("/auth/logout", authApi.LogOut)

	// 判断登录状态
	router.POST("/auth/is_login", authApi.IsLogin)

	// 第三方登录
	router.GET("/auth/github", authApi.AuthGithub)
	router.GET("/auth/github_callback", authApi.AuthGithubCallback)
}
