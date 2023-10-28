package middlewares

import (
	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

// Auth 鉴权接口

// registerAuthHandler is used to register auth handler
func registerAuthHandler() gin.HandlerFunc {
	// 先判断请求头是否为空，为空则为游客状态
	return func(c *gin.Context) {
		token := ""
		token = c.DefaultQuery("token", "")
		if token == "" {
			token = c.PostForm("token")
		}

		if token == "" {
			c.Next()
			return
		}

		// 判断token是否过期
		timeOut, err := utils.ValidToken(token)
		if timeOut || err != nil {
			// token过期或者解析token发生错误
			global.Logger.Errorf("token过期或者解析token发生错误!")
			c.Redirect(200, "/")
			return
		}

		// token没有过期,返回用户信息
		username, _ := utils.GetUsernameFromToken(token)
		nickname, _ := utils.GetNicknameFromToken(token)
		role, _ := utils.GetRoleFromToken(token)
		id, _ := utils.GetUserIdFromToken(token)
		c.Set("UserName", username)
		c.Set("NickName", nickname)
		c.Set("UserRole", role)
		c.Set("UserId", id)
		c.Next()

	}
}
