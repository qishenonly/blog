package models

import "time"

// UserCollectModel 用户收藏模型
type UserCollectModel struct {
	// User ID
	UserID uint `gorm:"type:int(11);not null;default:0;index:user_id" json:"user_id"`

	// User
	User UserModel `gorm:"foreignkey:UserID;association_foreignkey:ID;" json:"user"`

	// Article ID
	ArticleID uint `gorm:"type:int(11);not null;default:0;index:article_id" json:"article_id"`

	// Article
	Article ArticleModel `gorm:"foreignkey:ArticleID;association_foreignkey:ID;" json:"article"`

	// CreatedAt
	CreatedAt time.Time
}
