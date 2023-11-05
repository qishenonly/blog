package routers

import "github.com/qishenonly/blog/api"

func (router RouterGroup) InitImageRouter() {
	imageApi := api.NewApiGroup().ImageApi

	// 获取图片
	router.GET("/image/get_upload_image/:img_name", imageApi.GetUploadImage)
}
