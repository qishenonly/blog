package home

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func (hh *HomeApi) Home(c *gin.Context) {
	// 查询文章，按点赞数倒序排列，十条数据为一页
	utils.NewSuccessResponse("首页", c)

}
