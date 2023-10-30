package user

import (
	"github.com/qishenonly/blog/global"
	"github.com/qishenonly/blog/models"
	"github.com/qishenonly/blog/utils"

	"github.com/gin-gonic/gin"
)

type RequestToken struct {
	Token string `json:"token" binding:"required"`
}

type VisitorResponse struct {
	IsLogin bool   `json:"is_login"`
	Avatar  string `json:"avatar"`
	Motto   string `json:"motto"`
	Msg     string `json:"msg"`
}

type LoginUserResponse struct {
	IsLogin bool   `json:"is_login"`
	Avatar  string `json:"avatar"`
	Motto   string `json:"motto"`
	Msg     string `json:"msg"`
	QQ      string `json:"qq"`
	Github  string `json:"github"`
	Gitee   string `json:"gitee"`
	Zhihu   string `json:"zhihu"`
	CSDN    string `json:"csdn"`
}

func (ua *UserApi) GetUserInfo(c *gin.Context) {
	// 游客状态未登录
	var request RequestToken
	if err := c.ShouldBindJSON(&request); err != nil {
		global.Logger.Info("---", request.Token)
		visitorResponse := VisitorResponse{
			IsLogin: false,
			Avatar:  "https://picsum.photos/200/200",
			Motto:   utils.GenerateMotto(),
			Msg:     "用户未登录，当前为游客模式！",
		}
		utils.NewFailResponse(visitorResponse, c)
		return
	}

	// 校验 token
	ok, err := utils.ValidToken(request.Token)
	if err != nil && !ok {
		visitorResponse := VisitorResponse{
			IsLogin: false,
			Avatar:  "https://picsum.photos/200/200",
			Motto:   utils.GenerateMotto(),
			Msg:     "登录状态已过期，请重新登录！",
		}
		utils.NewFailResponse(visitorResponse, c)
		return
	}

	// 获取id
	id, err := utils.GetUserIdFromToken(request.Token)
	if err != nil {
		visitorResponse := VisitorResponse{
			IsLogin: false,
			Avatar:  "https://picsum.photos/200/200",
			Motto:   utils.GenerateMotto(),
			Msg:     "登录状态已过期，请重新登录！",
		}
		utils.NewFailResponse(visitorResponse, c)
		return
	}

	// 获取用户信息
	var user models.UserModel
	if err = global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		visitorResponse := VisitorResponse{
			IsLogin: false,
			Avatar:  "https://picsum.photos/200/200",
			Motto:   utils.GenerateMotto(),
			Msg:     "登录状态已过期，请重新登录！",
		}
		utils.NewFailResponse(visitorResponse, c)
		return
	}

	// 返回用户信息
	loginUserResponse := LoginUserResponse{
		IsLogin: true,
		Avatar:  user.Avatar,
		Motto:   user.Motto,
		Msg:     "获取用户信息成功！",
		QQ:      user.QQ,
		Github:  user.Github,
		Gitee:   user.Gitee,
		Zhihu:   user.Zhihu,
		CSDN:    user.CSDN,
	}

	utils.NewSuccessResponse(loginUserResponse, c)
}