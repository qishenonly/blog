package routers

import "github.com/qishenonly/blog/api"

func (router RouterGroup) InitUploadRouter() {
	uploadApi := api.NewApiGroup().UploadApi

	// 上传文章图片
	router.POST("/upload/article_image", uploadApi.UploadArticleImage)

	// 上传文章封面
	router.POST("/upload/article_cover", uploadApi.UploadArticleCoverImage)
}
