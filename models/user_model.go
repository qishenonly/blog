package models

import _type "blog/models/type"

// UserModel 是一个 UserModel 类型的结构体，用来映射数据库中的 users 表
type UserModel struct {
	Model

	// 昵称
	Nickname string `gorm:"type:varchar(20);not null;default:'';comment:'昵称'" json:"nickname"`

	// 用户名
	Username string `gorm:"type:varchar(20);not null;default:'';comment:'用户名'" json:"username"`

	// 密码
	Password string `gorm:"type:varchar(20);not null;default:'';comment:'密码'" json:"password"`

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
