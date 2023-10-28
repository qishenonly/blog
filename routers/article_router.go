package routers

import "blog/api"

func (router RouterGroup) InitArticleRouter() {
	articleApi := api.NewApiGroup().ArticleApi

	// 生成假数据
	router.GET("/article/faker", articleApi.Faker)

	// 文章列表
	router.POST("/article/list", articleApi.GetArticleList)
}
