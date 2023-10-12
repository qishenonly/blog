package auth

import (
	"blog/global"
	"blog/models"
	_type "blog/models/type"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

// Register godoc
// @Summary 注册
// @Description 注册
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Param data body Register true "data"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /auth/register [post]
func (ra *AuthApi) Register(c *gin.Context) {
	var request Register
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	if request.Username == "" || request.Password == "" || request.Email == "" || request.Code == "" {
		global.Logger.Error("参数错误", request.Email, request.Username, request.Password, request.Code)
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 判断验证码是否正确
	registerCode, err := global.Cache.Get([]byte("register_code_" + request.Code))
	if err != nil {
		global.Logger.Error("获取验证码失败: ", err)
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}
	if string(registerCode) != request.Code {
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}

	// 判断邮箱是否存在
	var user models.UserModel
	err = global.DB.Where("email = ?", request.Email).First(&user).Error
	if err == nil {
		utils.NewFailResponse("邮箱已存在", c)
		return
	}

	// 判断用户名是否存在
	err = global.DB.Where("username = ?", request.Username).First(&user).Error
	if err == nil {
		utils.NewFailResponse("用户名已存在", c)
		return
	}

	// 密码加密
	request.Password, err = utils.EncryptPassword(request.Password)
	if err != nil {
		global.Logger.Error("密码加密失败: ", err)
		utils.NewFailResponse("注册失败", c)
		return
	}

	// 添加用户
	err = global.DB.Create(&models.UserModel{
		Username:       request.Username,
		Password:       request.Password,
		Email:          request.Email,
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
