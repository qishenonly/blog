package article

import (
	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

type WriteArticleInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CoverPath   string `json:"cover_path"`
	CategoryID  uint   `json:"category_id"`
}

func (aa *ArticleApi) WriteArticle(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if _, err := utils.ParseToken(token); err != nil {
		utils.NewFailResponse("请先登录！", c)
		return
	}

	var articleInfo WriteArticleInfo
	if err := c.ShouldBindJSON(&articleInfo); err != nil {
		global.Logger.Error("参数错误: ", err)
		utils.NewFailResponse("参数错误", c)
		return
	}

	id, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Error("获取用户ID失败: ", err)
		utils.NewFailResponse("获取用户ID失败", c)
		return
	}

	article := models.ArticleModel{
		Title:        articleInfo.Title,
		Introduction: articleInfo.Description,
		Content:      articleInfo.Content,
		BannerPath:   articleInfo.CoverPath,
		CategoryID:   articleInfo.CategoryID,
		AuthorID:     uint(id),
	}

	if err := global.DB.Create(&article).Error; err != nil {
		global.Logger.Error("创建文章失败: ", err)
		utils.NewFailResponse("创建文章失败", c)
		return
	}

	utils.NewSuccessResponse(struct {
		ArticleID uint   `json:"article_id"`
		Msg       string `json:"msg"`
	}{
		ArticleID: article.ID,
		Msg:       "创建文章成功",
	}, c)
}
