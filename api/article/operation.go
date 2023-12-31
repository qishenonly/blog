package article

import (
	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

type ArticleRequest struct {
	ArticleID uint `json:"article_id"`
}

type ArticleResponse struct {
	Msg     string `json:"msg"`
	LikeNum uint   `json:"like_num"`
}

// Upvote godoc
//
//	@Summary		点赞
//	@Description	点赞
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			article_id	body		int		true	"文章ID"
//	@Success		200	{string}	string	"{"success":true,"data":{},"msg":"点赞成功"}"
//	@Router			/article/upvote [post]
func (aa *ArticleApi) Upvote(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if ok, _ := utils.ValidToken(token); ok {
		utils.NewFailResponse("请先登录", c)
		return
	}

	var request ArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Errorf("参数错误: %v", err)
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 判断是否已经点赞
	userID, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Errorf("获取用户ID失败: %v", err)
		utils.NewFailResponse("获取用户ID失败", c)
		return
	}

	// 判断文章是否存在
	var article models.ArticleModel
	if err := global.DB.Where("id = ?", request.ArticleID).First(&article).Error; err != nil {
		global.Logger.Errorf("获取文章失败: %v", err)
		utils.NewFailResponse("获取文章失败", c)
		return
	}

	// 判断用户是否存在
	var user models.UserModel
	if err := global.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		global.Logger.Errorf("获取用户失败: %v", err)
		utils.NewFailResponse("获取用户失败", c)
		return
	}

	// 判断是否已经点赞
	var likeAritcles []models.ArticleModel
	if err = global.DB.Model(&user).Association("LikeArticles").Find(&likeAritcles); err != nil {
		global.Logger.Errorf("获取点赞列表失败: %v", err)
		utils.NewFailResponse("获取点赞列表失败", c)
		return
	}
	for _, v := range likeAritcles {
		if v.ID == article.ID {
			utils.NewFailResponse("已经点赞", c)
			return
		}
	}

	// 点赞
	if err = global.DB.Model(&user).Association("LikeArticles").Append(&article); err != nil {
		global.Logger.Errorf("点赞失败: %v", err)
		utils.NewFailResponse("点赞失败", c)
		return
	}

	// 更新点赞数
	article.LikeNum++
	if err = global.DB.Model(&article).Update("like_num", article.LikeNum).Error; err != nil {
		global.Logger.Errorf("点赞失败: %v", err)
		utils.NewFailResponse("点赞失败", c)
		return
	}

	response := ArticleResponse{
		Msg:     "点赞成功",
		LikeNum: uint(article.LikeNum),
	}

	utils.NewSuccessResponse(response, c)

}

// CancelUpvote godoc
//
//	@Summary		取消点赞
//	@Description	取消点赞
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			article_id	body		int		true	"文章ID"
//	@Success		200	{string}	string	"{"success":true,"data":{},"msg":"取消点赞成功"}"
//	@Router			/article/cancel_upvote [post]
func (aa *ArticleApi) CancelUpvote(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if ok, _ := utils.ValidToken(token); ok {
		utils.NewFailResponse("请先登录", c)
		return
	}

	var request ArticleRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Errorf("参数错误: %v", err)
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 判断是否已经点赞
	userID, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Errorf("获取用户ID失败: %v", err)
		utils.NewFailResponse("获取用户ID失败", c)
		return
	}

	// 判断文章是否存在
	var article models.ArticleModel
	if err = global.DB.Where("id = ?", request.ArticleID).First(&article).Error; err != nil {
		global.Logger.Errorf("获取文章失败: %v", err)
		utils.NewFailResponse("获取文章失败", c)
		return
	}

	// 判断用户是否存在
	var user models.UserModel
	if err = global.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		global.Logger.Errorf("获取用户失败: %v", err)
		utils.NewFailResponse("获取用户失败", c)
		return
	}

	// 判断是否已经点赞
	var likeAritcles []models.ArticleModel
	if err = global.DB.Model(&user).Association("LikeArticles").Find(&likeAritcles); err != nil {
		global.Logger.Errorf("获取点赞列表失败: %v", err)
		utils.NewFailResponse("获取点赞列表失败", c)
		return
	}
	var flag bool
	for _, v := range likeAritcles {
		if v.ID == article.ID {
			flag = true
			break
		}
	}
	if !flag {
		utils.NewFailResponse("还没有点赞", c)
		return
	}

	// 取消点赞
	if err = global.DB.Model(&user).Association("LikeArticles").Delete(&article); err != nil {
		global.Logger.Errorf("取消点赞失败: %v", err)
		utils.NewFailResponse("取消点赞失败", c)
		return
	}

	// 更新点赞数
	article.LikeNum--
	if err = global.DB.Model(&article).Update("like_num", article.LikeNum).Error; err != nil {
		global.Logger.Errorf("取消点赞失败: %v", err)
		utils.NewFailResponse("取消点赞失败", c)
		return
	}

	response := ArticleResponse{
		Msg:     "取消点赞成功",
		LikeNum: uint(article.LikeNum),
	}

	utils.NewSuccessResponse(response, c)

}
