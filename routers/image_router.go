package routers

import "github.com/qishenonly/blog/api"

func (router RouterGroup) InitImageRouter() {
	imageApi := api.NewApiGroup().ImageApi

	// 获取文章图片
	router.GET("/image/article/:img_name", imageApi.GetUploadImage)

}
