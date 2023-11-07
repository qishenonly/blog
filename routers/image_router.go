package routers

import "github.com/qishenonly/blog/api"

func (router RouterGroup) InitImageRouter() {
	imageApi := api.NewApiGroup().ImageApi

	// 获取文章图片
	router.GET("/image/article/:img_name", imageApi.GetUploadArticleImage)

	// 获取文章封面
	router.GET("/image/article_cover/:img_name", imageApi.GetUploadArticleCover)
}
