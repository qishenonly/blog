package home

import (
	"blog/global"
	"blog/models"
	"github.com/gin-gonic/gin"
)

func (hh *HomeApi) Home(c *gin.Context) {
	// 查询文章，按点赞数倒序排列，十条数据为一页
	//articles, err := GetArticles(1, 10)
	//if err != nil {
	//	global.Logger.Error("获取文章失败: ", err)
	//	utils.NewFailResponse("获取文章失败", c)
	//	return
	//}

}

type ArticleInfo struct {
	ArticleID   uint   `json:"article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PageView    uint   `json:"page_view"`
	LikeNum     uint   `json:"like_num"`
	BannerPath  string `json:"banner_path"`
	CreateAt    string `json:"create_at"`
}

func (ah *HomeApi) ArticleList(c *gin.Context) {
	// 查询文章，按点赞数倒序排列，十条数据为一页
	//articles, err := GetArticles(1, 10)
	//if err != nil {
	//	global.Logger.Error("获取文章失败: ", err)
	//	utils.NewFailResponse("获取文章失败", c)
	//	return
	//}

}

// GetArticles 获取文章
func GetArticles(page, pageSize int) ([]models.ArticleModel, error) {
	var articles []models.ArticleModel

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询文章数据并限制结果数量
	err := global.DB.Limit(pageSize).Offset(offset).Find(&articles).Error
	if err != nil {
		global.Logger.Error("err: ", err)
		return nil, err
	}

	return articles, nil
}
