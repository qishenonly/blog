package auth

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func (va *AuthApi) Verify(c *gin.Context) {
	email := c.Param("email")

	// 判断邮箱是否存在
	var user models.UserModel
	err := global.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		utils.NewFailResponse("该用户不存在！", c)
		return
	}

	// 判断token是否正确
	token := c.Param("token")
	if token != user.ActivationToken {
		utils.NewFailResponse("激活失败！激活码不正确！", c)
		return
	}

	// 判断用户是否已激活
	if user.Activated && token == user.ActivationToken {
		utils.NewFailResponse("该用户已激活！", c)
		return
	}

	// 激活用户
	err = global.DB.Model(&user).Update("activated", true).Error
	if err != nil {
		global.Logger.Error("激活失败: ", err)
		utils.NewFailResponse("激活失败！", c)
		return
	}

	utils.NewSuccessResponse("激活成功", c)
}
