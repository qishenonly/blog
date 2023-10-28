package routers

import "github.com/qishenonly/blog/api"

func (router RouterGroup) InitHomeRouter() {
	homeApi := api.NewApiGroup().HomeApi

	// 首页
	router.GET("/home/", homeApi.Home)

}
