package image

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

// UploadArticleImage godoc
// @Summary 上传文章图片
// @Description 上传文章图片
// @Tags Image
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /image/upload_article_image [post]
func (ia *ImageApi) UploadArticleImage(c *gin.Context) {
	file, err := c.FormFile("file") // "file" 对应表单中的文件字段名
	if err != nil {
		global.Logger.Error("获取文件失败", err)
		utils.NewFailResponse("获取文件失败", c)
		return
	}

	contentType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		global.Logger.Error("只能上传图片")
		utils.NewFailResponse("只能上传图片", c)
		return
	}

	// 读取上传的文件数据
	fileData, err := file.Open()
	if err != nil {
		global.Logger.Error("读取文件失败", err)
		utils.NewFailResponse("读取文件失败", c)
		return
	}
	defer fileData.Close()

	data, err := ioutil.ReadAll(fileData)
	if err != nil {
		global.Logger.Error("读取文件失败", err)
		utils.NewFailResponse("读取文件失败", c)
		return
	}

	// 保存文件为图片
	fileName := utils.RandString(10) + "_" + file.Filename
	destinationPath := "upload/image/article/" + fileName
	err = os.WriteFile(destinationPath, data, 0644)
	if err != nil {
		global.Logger.Error("保存文件失败", err)
		utils.NewFailResponse("保存文件失败", c)
		return
	}

	utils.NewSuccessResponse(struct {
		Path string `json:"path"`
		Msg  string `json:"msg"`
	}{
		Path: global.Config.UploadImg.GetAddr() + "/api/image/get_upload_image/" + fileName,
		Msg:  "上传成功",
	}, c)
}
