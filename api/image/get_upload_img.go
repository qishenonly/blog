package image

import (
	"github.com/gin-gonic/gin"
)

// GetUploadArticleImage godoc
//
//	@Summary		获取图片
//	@Description	获取图片
//	@Tags			图片
//	@Accept			json
//	@Produce		json
//	@Param			img_name	path		string	true	"图片名称"
//	@Success		200			{object}	string	"图片"
//	@Router			/image/article/{img_name} [get]
func (ia *ImageApi) GetUploadArticleImage(c *gin.Context) {
	imgName := c.Param("img_name")
	c.File("upload/image/article/" + imgName)
}

func (ia *ImageApi) GetUploadArticleCover(c *gin.Context) {
	imgName := c.Param("img_name")
	c.File("upload/image/article_cover/" + imgName)
}
