package models

type CommentModel struct {
	Model

	// 评论内容
	Content string `gorm:"type:varchar(1000);default:''" json:"content"`

	// 评论者
	Commenter UserModel `gorm:"foreignkey:CommenterID;association_foreignkey:ID;" json:"commenter"`

	// 评论者 ID
	CommenterID uint `gorm:"not null;default:0;comment:'评论者 ID'" json:"commenter_id"`

	// 文章
	Article ArticleModel `gorm:"foreignkey:ArticleID;association_foreignkey:ID;" json:"article"`

	// 文章 ID
	ArticleID uint `gorm:"not null;default:0;comment:'文章 ID'" json:"article_id"`

	// 父评论
	ParentComment *CommentModel `gorm:"foreignkey:ParentCommentID;" json:"parent_comment"`

	// 父评论 ID
	ParentCommentID uint `gorm:"not null;default:0;comment:'父评论 ID'" json:"parent_comment_id"`

	// 子评论
	ChildComments []*CommentModel `gorm:"foreignkey:ParentCommentID;" json:"child_comments"`

	// 点赞数
	LikeNum int `gorm:"not null;default:0;comment:'点赞数'" json:"like_num"`

	// 踩数
	DislikeNum int `gorm:"not null;default:0;comment:'踩数'" json:"dislike_num"`
}
