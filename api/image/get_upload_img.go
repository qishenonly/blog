package image

import (
	"github.com/gin-gonic/gin"
)

// GetUploadImage godoc
//
//	@Summary		获取图片
//	@Description	获取图片
//	@Tags			图片
//	@Accept			json
//	@Produce		json
//	@Param			img_name	path		string	true	"图片名称"
//	@Success		200			{object}	string	"图片"
//	@Router			/image/get_upload_image/{img_name} [get]
func (ia *ImageApi) GetUploadImage(c *gin.Context) {
	imgName := c.Param("img_name")
	c.File("upload/image/article/" + imgName)
}
