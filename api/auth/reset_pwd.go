package auth

import (
	"blog/api/code"
	"blog/global"
	"blog/models"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type ResetPwd struct {
	NewPassword string `json:"new_password" binding:"required"`
	Code        string `json:"code" binding:"required"`
}

type ResetTokenActivatedData struct {
	Msg       string `json:"msg"`
	Activated bool   `json:"activated"`
}

type ResetFailRequestData struct {
	Msg        string          `json:"msg"`
	CustomCode code.CustomCode `json:"custom_code"`
}

// ResetPwd godoc
// @Summary 重置密码
// @Description 重置密码
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Param email path string true "email"
// @Param new_password body string true "new_password"
// @Param code body string true "code"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"重置密码成功"}"
// @Router /auth/reset_pwd/{email} [post]
func (ra *AuthApi) ResetPwd(c *gin.Context) {
	email := c.Param("email")
	if email == "" {
		response := ResetFailRequestData{
			Msg:        "获取参数失败",
			CustomCode: code.ParamError,
		}
		utils.NewFailResponse(response, c)
		return
	}

	// 判断邮箱是否存在
	var user models.UserModel
	if err := global.DB.Where("email = ?", email).First(&user).Error; err != nil {
		utils.NewFailResponse("该用户不存在！", c)
		return
	}

	// 没有重置过密码
	if user.ResetPasswordToken == "" {
		resetToken := utils.RandString(50)
		resetTokenData := utils.VerifyTokenData{
			Token: resetToken,
			Email: email,
		}
		// 发送邮件
		if err := utils.SendWithTemplate("reset_pwd_email",
			resetTokenData, email, "用户 "+user.Username+" 重置密码操作！"); err != nil {
			global.Logger.Error("发送邮件失败: ", err)
			utils.NewFailResponse("发送邮件失败！", c)
			return
		}

		// 保存重置密码token
		if err := global.DB.Model(&user).Update("reset_password_token",
			resetToken).Error; err != nil {
			global.Logger.Error("保存重置密码token失败: ", err)
			utils.NewFailResponse("重置密码失败！", c)
			return
		}

		// 保存重置密码token过期时间
		expireTime := time.Now().Add(time.Hour * 24).UnixNano()
		if err := global.DB.Model(&user).Update("reset_password_token_expired_at",
			expireTime).Error; err != nil {
			global.Logger.Error("保存重置密码token过期时间失败: ", err)
			utils.NewFailResponse("重置密码失败！", c)
		}

		data := ResetTokenActivatedData{
			Msg:       "重置密码激活链接已发送，请查收邮件激活！",
			Activated: false,
		}
		utils.NewSuccessResponse(data, c)
		return

	}

	// 重置过密码
	if user.ResetPasswordToken != "" {
		// 判断重置密码token是否过期
		if user.ResetPasswordTokenExpiredAt < time.Now().UnixNano() {
			resetToken := utils.RandString(50)
			resetTokenData := utils.VerifyTokenData{
				Token: resetToken,
				Email: email,
			}
			// 发送邮件
			if err := utils.SendWithTemplate("reset_pwd_email",
				resetTokenData, email, "用户 "+user.Username+" 重置密码操作！"); err != nil {
				global.Logger.Error("发送邮件失败: ", err)
				utils.NewFailResponse("发送邮件失败！", c)
				return
			}

			// 保存重置密码token
			if err := global.DB.Model(&user).Update("reset_password_token",
				resetToken).Error; err != nil {
				global.Logger.Error("保存重置密码token失败: ", err)
				utils.NewFailResponse("重置密码失败！", c)
				return
			}

			// 保存重置密码token过期时间
			expireTime := time.Now().Add(time.Hour * 24).UnixNano()
			if err := global.DB.Model(&user).Update("reset_password_token_expired_at",
				expireTime).Error; err != nil {
				global.Logger.Error("保存重置密码token过期时间失败: ", err)
				utils.NewFailResponse("重置密码失败！", c)
			}

			data := ResetTokenActivatedData{
				Msg:       "重置密码激活链接已发送，请查收邮件激活！",
				Activated: false,
			}
			utils.NewSuccessResponse(data, c)
			return
		}
	}

	// 判断是否激活
	if !user.ResetPasswordTokenUsed {
		data := ResetTokenActivatedData{
			Msg:       "重置密码激活链接未激活，请查收邮件激活！",
			Activated: false,
		}
		utils.NewSuccessResponse(data, c)
		return
	}

	// 获取参数
	var request ResetPwd
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		response := ResetFailRequestData{
			Msg:        "获取参数失败",
			CustomCode: code.ParamError,
		}
		utils.NewFailResponse(response, c)
		return
	}

	// 判断参数是否为空
	if request.NewPassword == "" || request.Code == "" {
		utils.NewFailResponse("参数不能为空", c)
		return
	}

	// 判断验证码是否正确
	resetPwdCode, err := global.Cache.Get([]byte("reset_pwd_code_" + request.Code))
	if err != nil {
		global.Logger.Error("获取验证码失败: ", err)
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}
	if string(resetPwdCode) != request.Code {
		global.Logger.Error("验证码错误: ", err)
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}

	// 判断新密码是否与旧密码相同
	newEncryptPwd, err := utils.EncryptPassword(request.NewPassword)
	if err != nil {
		global.Logger.Error("加密密码失败: ", err)
		utils.NewFailResponse("重置密码失败！", c)
		return
	}
	if err = utils.ComparePasswords(user.Password, newEncryptPwd); err == nil {
		utils.NewFailResponse("新密码不能与旧密码相同！", c)
		return
	}

	// 重置密码
	if err = global.DB.Model(&user).Update("password", newEncryptPwd).Error; err != nil {
		global.Logger.Error("重置密码失败: ", err)
		utils.NewFailResponse("重置密码失败！", c)
		return
	}

	// 删除验证码
	if err = global.Cache.Delete([]byte("reset_pwd_code_" + request.Code)); err != nil {
		global.Logger.Error("重置密码验证码删除失败: ", err)
		utils.NewSuccessResponse("重置密码成功", c)
		return
	}

	utils.NewSuccessResponse("重置密码成功", c)

}

// ResetPwdCode godoc
// @Summary 重置密码验证码
// @Description 重置密码验证码
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"重置密码验证码获取成功"}"
// @Router /auth/reset_pwd_code [get]
func (ra *AuthApi) ResetPwdCode(c *gin.Context) {
	code := utils.RandCode(4)
	resetPwdCode := []byte("reset_pwd_code_" + code)
	err := global.Cache.Put(resetPwdCode, []byte(code))
	if err != nil {
		global.Logger.Error("重置密码验证码缓存失败: ", err)
		utils.NewFailResponse("重置密码验证码获取失败！", c)
		return
	}
	global.Logger.Println("重置密码验证码：", code)
	utils.NewCodeResponse(code, "重置密码验证码！请在5分钟内使用！", c)
}
