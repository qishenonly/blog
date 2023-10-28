package article

import (
	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Title       string `fake:"{sentence:6}"`
	Description string `fake:"{sentence:15}"`
	Content     string `fake:"{paragraph:20}"`
	PageView    int    `fake:"{number:1,500}"`
	LikeNum     int    `fake:"{number:1,500}"`
	CommentNum  int    `fake:"{number:1,500}"`
	CollectNum  int    `fake:"{number:1,500}"`
}

// Faker godoc
// @Summary 生成文章数据
// @Description 生成文章数据
// @Tags Article
// @Accept json
// @Produce json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"生成文章数据成功"}"
// @Router /article/faker [post]
func (aa *ArticleApi) Faker(c *gin.Context) {
	var article Article

	// 生成标题，限制为 50 个汉字以内
	if err := gofakeit.Struct(&article); err != nil {
		global.Logger.Error("生成文章数据失败: ", err)
		utils.NewFailResponse("生成文章数据失败", c)
	}

	if err := global.DB.Create(&models.ArticleModel{
		Title:        article.Title,
		Introduction: article.Description,
		Content:      article.Content,
		PageView:     article.PageView,
		LikeNum:      article.LikeNum,
		CommentNum:   article.CommentNum,
		CollectNum:   article.CollectNum,
		AuthorID:     41,
	}).Error; err != nil {
		global.Logger.Error("文章数据插入失败: ", err)
		utils.NewFailResponse("生成文章数据失败", c)
	}

	utils.NewSuccessResponse("生成文章数据成功", c)

}
