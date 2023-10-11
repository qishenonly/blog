package auth

import (
	"blog/global"
	"blog/models"
	_type "blog/models/type"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary 注册
// @Description 注册
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string true "邮箱"
// @Param code formData string true "验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /auth/register [post]
func (ra *AuthApi) Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	code := c.PostForm("code")
	email := c.PostForm("email")

	if username == "" || password == "" || email == "" || code == "" {
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 判断验证码是否正确
	registerCode, err := global.Cache.Get([]byte("register_code_" + code))
	if err != nil {
		global.Logger.Error("获取验证码失败: ", err)
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}
	if string(registerCode) != code {
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}

	// 判断邮箱是否存在
	var user models.UserModel
	err = global.DB.Where("email = ?", email).First(&user).Error
	if err == nil {
		utils.NewFailResponse("邮箱已存在", c)
		return
	}

	// 判断用户名是否存在
	err = global.DB.Where("username = ?", username).First(&user).Error
	if err == nil {
		utils.NewFailResponse("用户名已存在", c)
		return
	}

	// 密码加密
	password, err = utils.EncryptPassword(password)
	if err != nil {
		global.Logger.Error("密码加密失败: ", err)
		utils.NewFailResponse("注册失败", c)
		return
	}

	// 添加用户
	err = global.DB.Create(&models.UserModel{
		Username:       username,
		Password:       password,
		Email:          email,
		Role:           _type.RoleUser,
		RegisterSource: _type.RegisterFromEmail,
	}).Error
	if err != nil {
		global.Logger.Error("注册失败: ", err)
		utils.NewFailResponse("注册失败", c)
		return
	}

	utils.NewSuccessResponse("注册成功", c)

}

// RegisterCode godoc
// @Summary 注册验证码
// @Description 注册验证码
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册验证码！请在5分钟内使用！"}"
// @Router /auth/register/code [get]
func (ra *AuthApi) RegisterCode(c *gin.Context) {
	code := utils.RandCode(4)
	registerCode := []byte("register_code_" + code)
	err := global.Cache.Put(registerCode, []byte(code))
	if err != nil {
		global.Logger.Error("注册验证码缓存失败: ", err)
		return
	}
	global.Logger.Println("注册验证码：", code)
	utils.NewCodeResponse(code, "注册验证码！请在5分钟内使用！", c)
}
