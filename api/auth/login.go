package auth

import (
	"blog/global"
	"blog/models"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Login godoc
// @Summary 登录
// @Description 登录
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Param email formData string true "邮箱"
// @Param password formData string true "密码"
// @Param code formData string true "验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登录成功"}"
// @Router /auth/login [post]
func (la *AuthApi) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	code := c.PostForm("code")
	if email == "" || password == "" || code == "" {
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 判断验证码是否正确
	loginCode, err := global.Cache.Get([]byte("login_code_" + code))
	if err != nil {
		global.Logger.Error("获取验证码失败: ", err)
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}
	if string(loginCode) != code {
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}

	// 判断邮箱是否存在
	var user models.UserModel
	err = global.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		utils.NewFailResponse("该用户不存在！", c)
		return
	}

	// 判断密码是否正确
	err = utils.ComparePasswords(user.Password, password)
	if err != nil {
		utils.NewFailResponse("密码错误！", c)
		return
	}

	// 判断用户是否已激活
	if !user.Activated {
		utils.NewFailResponse("该用户未激活！", c)
		return
	}

	// 生成token
	payload := utils.JwtPayload{
		UserId:   int64(user.ID),
		Username: user.Username,
		Role:     strconv.Itoa(int(user.Role)),
	}
	token, err := utils.GenerateToken(payload)
	if err != nil {
		global.Logger.Error("生成token失败: ", err)
		utils.NewFailResponse("登录失败", c)
		return
	}

	// Cache
	err = global.Cache.Put([]byte("token_"+strconv.Itoa(int(user.ID))), []byte(token))
	if err != nil {
		global.Logger.Error("token存入Cache失败: ", err)
		utils.NewFailResponse("登录失败", c)
		return
	}

	// 构建返回数据
	response := struct {
		Token string `json:"token"`
		Msg   string `json:"msg"`
	}{
		Token: token,
		Msg:   "登录成功",
	}
	utils.NewSuccessResponse(response, c)
}

// LoginCode godoc
// @Summary 登录验证码
// @Description 登录验证码
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登录验证码"}"
// @Router /auth/login/code [get]
func (la *AuthApi) LoginCode(c *gin.Context) {
	code := utils.RandCode(4)
	loginCode := []byte("login_code_" + code)
	err := global.Cache.Put(loginCode, []byte(code))
	if err != nil {
		global.Logger.Error("登录验证码缓存失败: ", err)
		utils.NewFailResponse("登录验证码获取失败！", c)
		return
	}
	global.Logger.Println("登录验证码：", code)
	utils.NewCodeResponse(code, "登录验证码！请在5分钟内使用！", c)
}
