package setting

import "github.com/gin-gonic/gin"

func (sa *SettingApi) SettingInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "SettingInfoView",
	})
}
