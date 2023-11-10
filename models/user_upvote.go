package models

import "time"

// UserUpvoteModel 用户点赞模型
type UserUpvoteModel struct {
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
