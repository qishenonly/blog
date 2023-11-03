package routers

import "github.com/qishenonly/blog/api"

func (router RouterGroup) InitUserRouter() {
	userApi := api.NewApiGroup().UserApi

	// 首页获取用户信息
	router.GET("/user/info", userApi.GetUserInfo)

	// 修改密码
	router.POST("/user/update_password", userApi.UpdateUserPassword)

	// 修改邮箱
	router.POST("/user/update_email", userApi.UpdateUserEmail)

	// 修改格言
	router.POST("/user/update_motto", userApi.UpdateUserMotto)

	// 修改社交帐号
	router.POST("/user/update_account_qq", userApi.UpdateUserSocialAccountQQ)
	router.POST("/user/update_account_github", userApi.UpdateUserSocialAccountGitHub)
	router.POST("/user/update_account_gitee", userApi.UpdateUserSocialAccountGitee)
	router.POST("/user/update_account_csdn", userApi.UpdateUserSocialAccountCSDN)

}
