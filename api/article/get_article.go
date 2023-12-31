package article

import (
	"time"

	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

type ArticleInfo struct {
	ArticleID   uint          `json:"article_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Content     string        `json:"content"`
	Comments    []CommentInfo `json:"comments"`
	Category    string        `json:"category"`
	Author      string        `json:"author"`
	AuthorID    uint          `json:"author_id"`
	Source      string        `json:"source"`
	SourceUrl   string        `json:"source_url"`
	PageView    uint          `json:"page_view"`
	LikeNum     uint          `json:"like_num"`
	CommentNum  uint          `json:"comment_num"`
	BannerPath  string        `json:"banner_path"`
	CreateAt    string        `json:"create_at"`
}

type UserModel struct {
	ID       uint   `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

type CommentInfo struct {
	ID          uint      `json:"id"`
	CommenterID uint      `json:"commenter_id"`
	Commenter   UserModel `json:"commenter"`
	Content     string    `json:"content"`
	LikeNum     uint      `json:"like_num"`
	DislikeNum  uint      `json:"dislike_num"`
	CreatedAt   string    `json:"created_at"`
}

// GetArticle godoc
//
//	@Summary		获取文章
//	@Description	获取文章
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int		true	"文章ID"
//	@Success		200	{string}	string	"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router			/article/detail/{id} [get]
func (aa *ArticleApi) GetArticle(c *gin.Context) {
	id := c.Param("id")
	var article models.ArticleModel
	if err := global.DB.Where("id = ?", id).First(&article).Error; err != nil {
		global.Logger.Error("获取文章失败: ", err)
		utils.NewFailResponse("获取文章失败", c)
		return
	}

	// 获取评论
	commentInfo := make([]CommentInfo, 0)
	if err := global.DB.Model(&models.CommentModel{}).Where("article_id = ?",
		id).Find(&commentInfo).Error; err != nil {
		global.Logger.Error("获取评论失败: ", err)
		utils.NewFailResponse("获取评论失败", c)
		return
	}

	// 获取评论者信息
	for i := 0; i < len(commentInfo); i++ {
		var commenter models.UserModel
		if err := global.DB.Where("id = ?", commentInfo[i].CommenterID).First(&commenter).Error; err != nil {
			global.Logger.Error("获取评论者信息失败: ", err)
			utils.NewFailResponse("获取评论者信息失败", c)
			return
		}
		commentInfo[i].Commenter = UserModel{
			ID:       commenter.ID,
			Nickname: commenter.Nickname,
			Avatar:   commenter.Avatar,
		}
	}

	// 获取分类
	var category models.CategoryModel
	if err := global.DB.Where("id = ?", article.CategoryID).First(&category).Error; err != nil {
		global.Logger.Error("获取分类失败: ", err)
		utils.NewFailResponse("获取分类失败", c)
		return
	}

	// 浏览量+1
	article.PageView++
	if err := global.DB.Model(&article).Where("id = ?", id).Update("page_view", article.PageView).Error; err != nil {
		global.Logger.Error("更新浏览量失败: ", err)
		utils.NewFailResponse("更新浏览量失败", c)
		return
	}

	articleInfo := ArticleInfo{
		ArticleID:   article.ID,
		Title:       article.Title,
		Description: article.Introduction,
		Content:     article.Content,
		Category:    category.Name,
		Author:      article.Author.Nickname,
		AuthorID:    article.AuthorID,
		Comments:    commentInfo,
		Source:      article.Source,
		SourceUrl:   article.SourceUrl,
		PageView:    uint(article.PageView),
		LikeNum:     uint(article.LikeNum),
		CommentNum:  uint(article.CommentNum),
		BannerPath:  article.BannerPath,
		CreateAt:    time.Unix(article.CreatedAt, 0).Format("2006-01-02 15:04:05"),
	}

	utils.NewSuccessResponse(articleInfo, c)

}
