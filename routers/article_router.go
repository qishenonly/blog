package routers

import "github.com/qishenonly/blog/api"

func (router RouterGroup) InitArticleRouter() {
	articleApi := api.NewApiGroup().ArticleApi

	// 生成假数据
	router.GET("/article/faker", articleApi.Faker)

	// 文章列表
	router.POST("/article/list", articleApi.GetArticleList)

	// 文章详情
	router.GET("/article/detail/:id", articleApi.GetArticle)

	// 文章分类
	router.GET("/article/category", articleApi.GetArticleCategory)

	// 写文章
	router.POST("/article/write", articleApi.WriteArticle)

	// 点赞
	router.POST("/article/upvote", articleApi.Upvote)

	// 取消点赞
	router.POST("/article/cancel_upvote", articleApi.CancelUpvote)

}
