package auth

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func (va *AuthApi) Verify(c *gin.Context) {
	utils.NewSuccessResponse("验证成功", c)
}
