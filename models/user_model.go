package models

import (
	"time"

	"github.com/qishenonly/blog/global"
	_type "github.com/qishenonly/blog/models/type"
	"github.com/qishenonly/blog/utils"

	"gorm.io/gorm"
)

// UserModel 是一个 UserModel 类型的结构体，用来映射数据库中的 users 表
type UserModel struct {
	Model

	// 昵称
	Nickname string `gorm:"type:varchar(20);not null;default:'';comment:'昵称'" json:"nickname"`

	// 用户名
	Username string `gorm:"type:varchar(20);not null;default:'';comment:'用户名'" json:"username"`

	// 密码
	Password string `gorm:"type:varchar(100);not null;default:'';comment:'密码'" json:"password"`

	// 头像
	Avatar string `gorm:"type:varchar(100);not null;default:'';comment:'头像'" json:"avatar"`

	// 邮箱
	Email string `gorm:"type:varchar(50);not null;default:'';comment:'邮箱'" json:"email"`

	// 邮箱是否验证
	EmailVerified bool `gorm:"not null;default:false;comment:'邮箱是否验证'" json:"email_verified"`

	// 手机号
	Phone string `gorm:"type:varchar(20);not null;default:'';comment:'手机号'" json:"phone"`

	// 手机号是否验证
	PhoneVerified bool `gorm:"not null;default:false;comment:'手机号是否验证'" json:"phone_verified"`

	// 其他平台的唯一标识
	OpenId string `gorm:"type:varchar(50);not null;default:'';comment:'其他平台的唯一标识'" json:"open_id"`

	// 注册来源
	RegisterSource _type.RegisterSource `gorm:"type:varchar(20);not null;default:1;comment:'注册来源'" json:"RegisterSource"`

	// 激活令牌
	ActivationToken string `gorm:"type:varchar(100);not null;default:'';comment:'激活令牌'" json:"activation_token"`

	// 激活令牌是否已经被激活
	Activated bool `gorm:"not null;default:false;comment:'激活令牌是否已经被激活'" json:"activated"`

	// 激活令牌过期时间
	ActivationTokenExpiredAt int64 `gorm:"not null;default:0;comment:'激活令牌过期时间'" json:"activation_token_expired_at"`

	// 重置密码令牌
	ResetPasswordToken string `gorm:"type:varchar(100);not null;default:'';comment:'重置密码令牌'" json:"reset_password_token"`

	// 重置密码令牌过期时间
	ResetPasswordTokenExpiredAt int64 `gorm:"not null;default:0;comment:'重置密码令牌过期时间'" json:"reset_password_token_expired_at"`

	// 重置密码令牌是否已经被使用
	ResetPasswordTokenUsed bool `gorm:"not null;default:false;comment:'重置密码令牌是否已经被使用'" json:"reset_password_token_used"`

	// 地址
	Address string `gorm:"type:varchar(100);not null;default:'';comment:'地址'" json:"address"`

	// IP
	Ip string `gorm:"type:varchar(20);not null;default:'';comment:'IP'" json:"ip"`

	// 权限   1--管理员  2--普通用户  3--游客  4--被封禁的用户
	Role _type.Role `gorm:"not null;default:3;comment:'权限'" json:"role"`

	// 状态   1--正常  2--被封禁
	Status int `gorm:"not null;default:1;comment:'状态'" json:"status"`

	// 性别   1--男  2--女  3--保密
	Sex int `gorm:"not null;default:1;comment:'性别'" json:"Sex"`

	// 发布的文章列表
	Articles []ArticleModel `gorm:"many2many:user_article;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:ID;jointable_foreignkey:ID;" json:"articles"`

	// 收藏的文章列表
	CollectArticles []ArticleModel `gorm:"many2many:user_collect;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:ID;jointable_foreignkey:ID;" json:"collect_articles"`

	// 关注的用户列表
	FollowUsers []UserModel `gorm:"many2many:user_follow_user;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:ID;jointable_foreignkey:ID;" json:"follow_users"`

	// 粉丝列表
	FansUsers []UserModel `gorm:"many2many:user_fans_user;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:ID;jointable_foreignkey:ID;" json:"fans_users"`
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.ActivationToken = utils.RandString(50)
	u.ActivationTokenExpiredAt = time.Now().Add(time.Hour * 24).UnixNano()
	return
}

func (u *UserModel) AfterCreate(tx *gorm.DB) (err error) {
	tokendata := utils.VerifyTokenData{
		Token: u.ActivationToken,
		Email: u.Email,
	}
	err = utils.SendWithTemplate("confirm_email", tokendata, u.Email, "用户 "+u.Username+" 激活帐号操作！")
	if err != nil {
		global.DB.Delete(u)
		return err
	}
	return
}
