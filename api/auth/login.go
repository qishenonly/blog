package auth

import (
	"strconv"

	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required"`
}

// Login godoc
// @Summary 登录
// @Description 登录
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Param email body string true "email"
// @Param password body string true "password"
// @Param code body string true "code"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登录成功"}"
// @Router /auth/login [post]
func (la *AuthApi) Login(c *gin.Context) {
	var request Login
	if err := c.ShouldBind(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}
	if request.Email == "" || request.Password == "" || request.Code == "" {
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 判断验证码是否正确
	loginCode, err := global.Cache.Get([]byte("login_code_" + request.Code))
	if err != nil {
		global.Logger.Error("获取验证码失败: ", err)
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}
	if string(loginCode) != request.Code {
		utils.NewFailResponse("验证码错误！请重试！", c)
		return
	}

	// 判断邮箱是否存在
	var user models.UserModel
	err = global.DB.Where("email = ?", request.Email).First(&user).Error
	if err != nil {
		utils.NewFailResponse("该用户不存在！", c)
		return
	}

	// 判断密码是否正确
	err = utils.ComparePasswords(user.Password, request.Password)
	if err != nil {
		utils.NewFailResponse("密码错误！", c)
		return
	}

	// 判断用户是否已激活
	if !user.Activated {
		tokendata := utils.VerifyTokenData{
			Token: user.ActivationToken,
			Email: user.Email,
		}
		if err = utils.SendWithTemplate("confirm_email", tokendata,
			user.Email, "用户 "+user.Username+" 激活帐号操作！"); err != nil {
			global.Logger.Error("发送邮件失败: ", err)
			utils.NewFailResponse("发送邮件失败！", c)
			return
		}
		utils.NewFailResponse("该用户未激活！激活链接邮件已发送！请注意查收！", c)
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

	// ip
	if err = global.DB.Model(&user).Update("ip", c.ClientIP()).Error; err != nil {
		global.Logger.Error("更新用户ip失败: ", err)
		return
	}

	// Cache
	err = global.Cache.Put([]byte("token_"+strconv.Itoa(int(user.ID))), []byte(token))
	if err != nil {
		global.Logger.Error("token存入Cache失败: ", err)
		utils.NewFailResponse("登录失败", c)
		return
	}

	// 删除验证码
	if err = global.Cache.Delete([]byte("login_code_" + request.Code)); err != nil {
		global.Logger.Error("登录验证码删除失败: ", err)
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

type Token struct {
	Token string `json:"token" binding:"required"`
}

type Response struct {
	IsLogin bool   `json:"is_login"`
	Msg     string `json:"msg"`
}

// IsLogin godoc
// @Summary 判断是否登录
// @Description 判断是否登录
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Param token body string true "token"
func (la *AuthApi) IsLogin(c *gin.Context) {
	var request Token
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断token是否过期
	ok, err := utils.ValidToken(request.Token)
	if err != nil && !ok {
		response := Response{
			IsLogin: false,
			Msg:     "token过期！",
		}
		global.Logger.Error("token验证失败: ", err)
		utils.NewFailResponse(response, c)
		return
	}

	response := Response{
		IsLogin: true,
		Msg:     "token校验成功！",
	}
	utils.NewSuccessResponse(response, c)
}

// LogOut godoc
// @Summary 退出登录
// @Description 退出登录
// @Tags Auth
// @Accept  application/json
// @Produce  application/json
// @Param token body string true "token"
func (la *AuthApi) LogOut(c *gin.Context) {
	var request Token
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断token是否过期
	ok, err := utils.ValidToken(request.Token)
	if err != nil && !ok {
		response := Response{
			IsLogin: false,
			Msg:     "token过期！",
		}
		global.Logger.Error("token验证失败: ", err)
		utils.NewFailResponse(response, c)
		return
	}

	// 获取token中的用户id
	id, err := utils.GetUserIdFromToken(request.Token)
	if err != nil {
		global.Logger.Error("获取用户id失败: ", err)
		utils.NewFailResponse("退出登录失败！", c)
		return
	}
	// 删除token
	err = global.Cache.Delete([]byte("token_" + strconv.Itoa(int(id))))
	if err != nil {
		global.Logger.Error("token删除失败: ", err)
		utils.NewFailResponse("退出登录失败！", c)
		return
	}

	response := Response{
		IsLogin: false,
		Msg:     "退出登录成功！",
	}
	utils.NewSuccessResponse(response, c)
}