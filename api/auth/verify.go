package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

// Verify godoc
//	@Summary		激活账号
//	@Description	激活账号
//	@Tags			Auth
//	@Accept			application/json
//	@Produce		application/json
//	@Param			token	path		string	true	"激活码"
//	@Param			email	path		string	true	"邮箱"
//	@Success		200		{string}	string	"{"success":true,"data":{},"msg":"激活成功"}"
//	@Router			/auth/verify/{token}/active/{email} [get]
func (va *AuthApi) Verify(c *gin.Context) {
	email := c.Param("email")

	// 判断邮箱是否存在
	var user models.UserModel
	err := global.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err = WriteActiveStatusToCache(email, 310, "该链接已失效！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		return
	}

	// 判断token是否过期
	currentTime := time.Now().UnixNano()
	expireTime := user.ActivationTokenExpiredAt
	if currentTime > expireTime {
		if err := WriteActiveStatusToCache(email, 310, "该链接已失效！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		return
	}

	// 判断token是否正确
	token := c.Param("token")
	if token != user.ActivationToken {
		if err = WriteActiveStatusToCache(email, 310, "该链接已失效！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		//utils.NewFailResponse("该链接已失效！", c)
		return
	}

	// 判断用户是否已激活
	if user.Activated && token == user.ActivationToken {
		if err = WriteActiveStatusToCache(email, 310, "该用户已激活！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		return
	}

	// 激活用户
	err = global.DB.Model(&user).Update("activated", true).Error
	if err != nil {
		if err = WriteActiveStatusToCache(email, 310, "该链接已失效！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		return
	}

	if err = WriteActiveStatusToCache(email, 312, "激活成功！！", "success",
		http.StatusFound, "http://localhost:8888/", c); err != nil {
		return
	}
}

// VerifyResetPwd godoc
//	@Summary		激活重置密码token
//	@Description	激活重置密码token
//	@Tags			Auth
//	@Accept			application/json
//	@Produce		application/json
//	@Param			token	path		string	true	"激活码"
//	@Param			email	path		string	true	"邮箱"
//	@Success		200		{string}	string	"{"success":true,"data":{},"msg":"激活成功"}"
//	@Router			/auth/verify/{token}/reset_pwd/{email} [get]
func (va *AuthApi) VerifyResetPwd(c *gin.Context) {
	email := c.Param("email")

	// 判断邮箱是否存在
	var user models.UserModel
	if err := global.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err = WriteActiveStatusToCache(email, 310, "该链接已失效！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		return
	}

	// 判断token是否过期
	currentTime := time.Now().UnixNano()
	expireTime := user.ResetPasswordTokenExpiredAt
	if currentTime > expireTime {
		if err := WriteActiveStatusToCache(email, 310, "该链接已失效！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		return
	}

	// 判断token是否正确
	// 判断token是否已激活
	token := c.Param("token")
	if token != user.ResetPasswordToken || user.ResetPasswordTokenUsed {
		if err := WriteActiveStatusToCache(email, 310, "该链接已失效！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		return
	}

	// 激活重置密码token
	if err := global.DB.Model(&user).
		Update("reset_password_token_used", true).Error; err != nil {
		global.Logger.Error("激活失败: ", err)
		if err = WriteActiveStatusToCache(email, 310, "该链接已失效！", "error",
			http.StatusFound, "http://localhost:8888/", c); err != nil {
			return
		}
		return
	}

	if err := WriteActiveStatusToCache(email, 312, "激活成功！！", "success",
		http.StatusFound, "http://localhost:8888/", c); err != nil {
		return
	}
}

type ActiveMsg struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

// WriteActiveStatusToCache 写入激活状态到cache
// email: 邮箱
// code: 返回状态码
// msg: 返回信息
// status: 激活状态
// HttpStatusFound: 重定向状态码
// location: 重定向地址
// c: gin.Context
func WriteActiveStatusToCache(email string, code int, msg string, status string,
	HttpStatusFound int, location string, c *gin.Context) error {
	activeMsg := ActiveMsg{
		Msg:  msg,
		Code: code,
	}
	jsonMsg, err := json.Marshal(activeMsg)
	if err != nil {
		return err
	}
	if err = global.Cache.Put([]byte("active_"+status+"_"+email), jsonMsg); err != nil {
		global.Logger.Error("写入cache失败: ", err)
		return err
	}
	utils.NewRedirectResponse(HttpStatusFound, location, c)
	return nil
}
