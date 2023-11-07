package article

import (
	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

type CategoryInfo struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetArticleCategory godoc
//
//	@Summary		获取文章分类
//	@Description	获取文章分类
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router			/article/category [get]
func (aa *ArticleApi) GetArticleCategory(c *gin.Context) {
	var category []models.CategoryModel
	if err := global.DB.Find(&category).Error; err != nil {
		global.Logger.Error("获取文章分类失败", err)
		utils.NewFailResponse("获取文章分类失败", c)
		return
	}

	var categoryList []CategoryInfo
	for _, v := range category {
		categoryInfo := CategoryInfo{
			ID:   v.ID,
			Name: v.Name,
		}
		categoryList = append(categoryList, categoryInfo)
	}

	utils.NewSuccessResponse(categoryList, c)
}
