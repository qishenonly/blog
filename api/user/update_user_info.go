package user

import (
	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UpdateEmailRequest 用来映射修改邮箱请求参数
type UpdateEmailRequest struct {
	Email string `json:"email" binding:"required"`
}

// UpdateUserEmail godoc
// @Summary 更新用户邮箱
// @Description 更新用户邮箱
// @Tags User
// @Accept  application/json
// @Produce  application/json
// @Param email body string true "email"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新用户邮箱成功"}"
// @Router /user/update_email [post]
func (ua *UserApi) UpdateUserEmail(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var request UpdateEmailRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断参数是否为空
	if token == "" || request.Email == "" {
		global.Logger.Error("参数错误")
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 校验 token
	ok, err := utils.ValidToken(token)
	if err != nil && !ok {
		global.Logger.Error("token校验失败！", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 获取用户id
	id, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Error("Token过期: ", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 根据用户id获取用户信息
	var user models.UserModel
	if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		global.Logger.Error("获取用户信息失败: ", err)
		utils.NewFailResponse("获取用户信息失败！", c)
		return
	}

	// 判断邮箱是否与原来的相同
	if user.Email == request.Email {
		global.Logger.Error("新邮箱与旧邮箱相同")
		utils.NewFailResponse("新邮箱与旧邮箱相同！", c)
		return
	}

	// 更新用户邮箱
	if err = updateUserField(global.DB, user, "email", request.Email); err != nil {
		global.Logger.Error("更新用户邮箱失败: ", err)
		utils.NewFailResponse("更新用户邮箱失败！", c)
		return
	}

	utils.NewSuccessResponse("更新用户邮箱成功！", c)
}

// UpdatePwdRequest 用来映射修改密码请求参数
type UpdatePwdRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

// UpdateUserPassword godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags User
// @Accept  application/json
// @Produce  application/json
// @Param new_password body string true "new_password"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新用户密码成功！"}"
// @Router /user/update_password [post]
func (ua *UserApi) UpdateUserPassword(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var request UpdatePwdRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断参数是否为空
	if token == "" || request.NewPassword == "" {
		global.Logger.Error("参数错误")
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 校验 token
	ok, err := utils.ValidToken(token)
	if err != nil && !ok {
		global.Logger.Error("token校验失败！", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 获取用户id
	id, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Error("Token过期: ", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 根据用户id获取用户信息
	var user models.UserModel
	if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		global.Logger.Error("获取用户信息失败: ", err)
		utils.NewFailResponse("获取用户信息失败！", c)
		return
	}

	// 判断密码是否与原密码相同
	if err = utils.ComparePasswords(user.Password, request.NewPassword); err == nil {
		global.Logger.Error("新密码与原密码相同！")
		utils.NewFailResponse("新密码与原密码相同！", c)
		return
	}

	// 密码加密
	request.NewPassword, err = utils.EncryptPassword(request.NewPassword)
	if err != nil {
		global.Logger.Error("密码加密失败: ", err)
		utils.NewFailResponse("更新用户密码失败", c)
		return
	}

	// 更新用户密码
	if err = updateUserField(global.DB, user, "password", request.NewPassword); err != nil {
		global.Logger.Error("更新用户密码失败: ", err)
		utils.NewFailResponse("更新用户密码失败！", c)
		return
	}

	utils.NewSuccessResponse("更新用户密码成功！", c)

}

// UpdateMottoRequest 用来映射修改motto请求参数
type UpdateMottoRequest struct {
	Motto string `json:"motto" binding:"required"`
}

// UpdateUserMotto godoc
// @Summary 修改用户motto
// @Description 修改用户motto
// @Tags User
// @Accept  application/json
// @Produce  application/json
// @Param motto body string true "motto"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新用户motto成功！"}"
// @Router /user/update_motto [post]
func (ua *UserApi) UpdateUserMotto(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var request UpdateMottoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断参数是否为空
	if token == "" || request.Motto == "" {
		global.Logger.Error("参数错误")
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 校验 token
	ok, err := utils.ValidToken(token)
	if err != nil && !ok {
		global.Logger.Error("token校验失败！", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 获取用户id
	id, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Error("Token过期: ", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 根据用户id获取用户信息
	var user models.UserModel
	if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		global.Logger.Error("获取用户信息失败: ", err)
		utils.NewFailResponse("获取用户信息失败！", c)
		return
	}

	// 更新用户motto
	if err = updateUserField(global.DB, user, "motto", request.Motto); err != nil {
		global.Logger.Error("更新用户motto失败: ", err)
		utils.NewFailResponse("更新用户motto失败！", c)
		return
	}

	utils.NewSuccessResponse("更新用户motto成功！", c)
}

// UpdateSocialAccountQQRequest 用来映射修改社交账号请求参数
type UpdateSocialAccountQQRequest struct {
	QQ string `json:"qq" binding:"required"`
}

// UpdateUserSocialAccountQQ godoc
// @Summary 修改用户社交账号
// @Description 修改用户社交账号
// @Tags User
// @Accept  application/json
// @Produce  application/json
// @Param qq body string true "qq"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新用户社交账号成功！"}"
// @Router /user/update_social_account_qq [post]
func (ua *UserApi) UpdateUserSocialAccountQQ(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var request UpdateSocialAccountQQRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断参数是否为空
	if token == "" && request.QQ == "" {
		global.Logger.Error("参数错误")
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 校验 token
	ok, err := utils.ValidToken(token)
	if err != nil && !ok {
		global.Logger.Error("token校验失败！", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 获取用户id
	id, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Error("Token过期: ", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 根据用户id获取用户信息
	var user models.UserModel
	if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		global.Logger.Error("获取用户信息失败: ", err)
		utils.NewFailResponse("获取用户信息失败！", c)
		return
	}

	// 更新用户社交账号
	if err = updateUserField(global.DB, user, "qq", request.QQ); err != nil {
		global.Logger.Error("更新用户QQ失败: ", err)
		utils.NewFailResponse("更新用户QQ失败！", c)
		return
	}

	utils.NewSuccessResponse("更新用户QQ社交账号成功！", c)
}

// UpdateSocialAccountGithubRequest 用来映射修改社交账号请求参数
type UpdateSocialAccountGithubRequest struct {
	Github string `json:"github" binding:"required"`
}

// UpdateUserSocialAccountGitHub godoc
// @Summary 修改用户社交账号
// @Description 修改用户社交账号
// @Tags User
// @Accept  application/json
// @Produce  application/json
// @Param github body string true "github"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新用户社交账号成功！"}"
// @Router /user/update_social_account_github [post]
func (ua *UserApi) UpdateUserSocialAccountGitHub(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var request UpdateSocialAccountGithubRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断参数是否为空
	if token == "" && request.Github == "" {
		global.Logger.Error("参数错误")
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 校验 token
	ok, err := utils.ValidToken(token)
	if err != nil && !ok {
		global.Logger.Error("token校验失败！", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 获取用户id
	id, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Error("Token过期: ", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 根据用户id获取用户信息
	var user models.UserModel
	if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		global.Logger.Error("获取用户信息失败: ", err)
		utils.NewFailResponse("获取用户信息失败！", c)
		return
	}

	// 更新用户社交账号
	if err = updateUserField(global.DB, user, "github", request.Github); err != nil {
		global.Logger.Error("更新用户Github失败: ", err)
		utils.NewFailResponse("更新用户Github失败！", c)
		return
	}

	utils.NewSuccessResponse("更新用户GitHub社交账号成功！", c)
}

// UpdateSocialAccountGiteeRequest 用来映射修改社交账号请求参数
type UpdateSocialAccountGiteeRequest struct {
	Gitee string `json:"gitee" binding:"required"`
}

// UpdateUserSocialAccountGitee godoc
// @Summary 修改用户社交账号
// @Description 修改用户社交账号
// @Tags User
// @Accept  application/json
// @Produce  application/json
// @Param github body string true "gitee"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新用户社交账号成功！"}"
// @Router /user/update_social_account_gitee [post]
func (ua *UserApi) UpdateUserSocialAccountGitee(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var request UpdateSocialAccountGiteeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断参数是否为空
	if token == "" && request.Gitee == "" {
		global.Logger.Error("参数错误")
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 校验 token
	ok, err := utils.ValidToken(token)
	if err != nil && !ok {
		global.Logger.Error("token校验失败！", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 获取用户id
	id, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Error("Token过期: ", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 根据用户id获取用户信息
	var user models.UserModel
	if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		global.Logger.Error("获取用户信息失败: ", err)
		utils.NewFailResponse("获取用户信息失败！", c)
		return
	}

	// 更新用户社交账号
	if err = updateUserField(global.DB, user, "gitee", request.Gitee); err != nil {
		global.Logger.Error("更新用户Gitee失败: ", err)
		utils.NewFailResponse("更新用户Gitee失败！", c)
		return
	}

	utils.NewSuccessResponse("更新用户Gitee社交账号成功！", c)
}

// UpdateSocialAccountCSDNRequest 用来映射修改社交账号请求参数
type UpdateSocialAccountCSDNRequest struct {
	CSDN string `json:"csdn" binding:"required"`
}

// UpdateUserSocialAccountCSDN godoc
// @Summary 修改用户社交账号
// @Description 修改用户社交账号
// @Tags User
// @Accept  application/json
// @Produce  application/json
// @Param github body string true "csdn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新用户社交账号成功！"}"
// @Router /user/update_social_account_csdn [post]
func (ua *UserApi) UpdateUserSocialAccountCSDN(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var request UpdateSocialAccountCSDNRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Error("获取参数失败", err)
		utils.NewFailResponse("获取参数失败", c)
		return
	}

	// 判断参数是否为空
	if token == "" && request.CSDN == "" {
		global.Logger.Error("参数错误")
		utils.NewFailResponse("参数错误", c)
		return
	}

	// 校验 token
	ok, err := utils.ValidToken(token)
	if err != nil && !ok {
		global.Logger.Error("token校验失败！", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 获取用户id
	id, err := utils.GetUserIdFromToken(token)
	if err != nil {
		global.Logger.Error("Token过期: ", err)
		utils.NewFailResponse("Token过期！", c)
		return
	}

	// 根据用户id获取用户信息
	var user models.UserModel
	if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		global.Logger.Error("获取用户信息失败: ", err)
		utils.NewFailResponse("获取用户信息失败！", c)
		return
	}

	// 更新用户社交账号
	if err = updateUserField(global.DB, user, "csdn", request.CSDN); err != nil {
		global.Logger.Error("更新用户CSDN失败: ", err)
		utils.NewFailResponse("更新用户CSDN失败！", c)
		return
	}

	utils.NewSuccessResponse("更新用户CSDN社交账号成功！", c)
}

// 更新用户字段
func updateUserField(db *gorm.DB, user models.UserModel, field string, value string) error {
	if value != "" {
		if err := db.Model(&user).Update(field, value).Error; err != nil {
			return err
		}
	}
	return nil
}
