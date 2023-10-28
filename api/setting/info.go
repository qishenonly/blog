package setting

import (
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

func (sa *SettingApi) SettingInfoView(c *gin.Context) {
	utils.NewSuccessResponse("This is SettingInfoView!", c)
}
