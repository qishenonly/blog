package setting

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func (sa *SettingApi) SettingInfoView(c *gin.Context) {
	utils.NewSuccessResponse("This is SettingInfoView!", c)
}
