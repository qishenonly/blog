package article

import (
	"time"

	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

type ArticleInfos struct {
	ArticleID   uint   `json:"article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PageView    uint   `json:"page_view"`
	LikeNum     uint   `json:"like_num"`
	CommentNum  uint   `json:"comment_num"`
	BannerPath  string `json:"banner_path"`
	CreateAt    string `json:"create_at"`
}

type PageInfo struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"page_size" binding:"required"`
}

// GetArticleList godoc
//	@Summary		获取文章列表
//	@Description	获取文章列表
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			page		body		int		true	"页码"
//	@Param			page_size	body		int		true	"每页数量"
//	@Success		200			{string}	string	"{"success":true,"data":{},"msg":"获取成功"}"
//	@Router			/article/list [post]
func (aa *ArticleApi) GetArticleList(c *gin.Context) {
	// 查询文章，按点赞数倒序排列，十条数据为一页
	var request PageInfo
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败: ", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}
	articles, err := GetArticles(request.Page, request.PageSize)
	if err != nil {
		global.Logger.Error("获取文章失败: ", err)
		utils.NewFailResponse("获取文章失败", c)
		return
	}

	var total int64
	if err = global.DB.Model(&models.ArticleModel{}).Count(&total).Error; err != nil {
		global.Logger.Error("获取文章总数失败: ", err)
		utils.NewFailResponse("获取文章总数失败", c)
		return
	}

	// 判断是否还有更多数据
	var hasMore bool
	if total > int64(request.Page*request.PageSize) {
		hasMore = true
	} else {
		hasMore = false
	}

	response := struct {
		ArticleList []ArticleInfos `json:"article_list"`
		HasMore     bool           `json:"has_more"`
		CurrentPage int            `json:"current_page"`
	}{
		ArticleList: articles,
		HasMore:     hasMore,
		CurrentPage: request.Page,
	}

	utils.NewSuccessResponse(response, c)

}

// GetArticles 获取文章
func GetArticles(page, pageSize int) ([]ArticleInfos, error) {
	var articles []models.ArticleModel

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询文章数据并限制结果数量
	err := global.DB.Limit(pageSize).Offset(offset).Find(&articles).Error
	if err != nil {
		global.Logger.Error("err: ", err)
		return nil, err
	}

	var ArticleInfoss []ArticleInfos
	for _, article := range articles {
		ArticleInfoss = append(ArticleInfoss, ArticleInfos{
			ArticleID:   article.ID,
			Title:       article.Title,
			Description: article.Introduction,
			PageView:    uint(article.PageView),
			LikeNum:     uint(article.LikeNum),
			CommentNum:  uint(article.CommentNum),
			BannerPath:  article.BannerPath,
			CreateAt:    time.Unix(article.CreatedAt, 0).Format("2006-01-02 15:04:05"),
		})
	}

	return ArticleInfoss, nil
}
