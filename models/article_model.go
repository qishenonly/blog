package models

import _type "blog/models/type"

// ArticleModel 是一个 ArticleModel 类型的结构体，用来映射数据库中的 articles 表
type ArticleModel struct {
	Model

	// 标题
	Title string `gorm:"type:varchar(50);not null;default:'';comment:'标题'" json:"title"`

	// 简介
	Introduction string `gorm:"type:varchar(100);not null;default:'';comment:'简介'" json:"introduction"`

	// 内容
	Content string `gorm:"type:longtext;not null;comment:'内容'" json:"content"`

	// 浏览量
	PageView int `gorm:"not null;default:0;comment:'浏览量'" json:"page_view"`

	// 点赞数
	LikeNum int `gorm:"not null;default:0;comment:'点赞数'" json:"like_num"`

	// 收藏数
	CollectNum int `gorm:"not null;default:0;comment:'收藏数'" json:"collect_num"`

	// 评论数
	CommentNum int `gorm:"not null;default:0;comment:'评论数'" json:"comment_num"`

	// 评论列表
	Comments []CommentModel `gorm:"foreignkey:ID;association_foreignkey:ID;" json:"comments"`

	// 分类
	Category string `gorm:"type:varchar(20);not null;default:'';comment:'分类（标签）'" json:"category"`

	// 文章作者
	Author UserModel `gorm:"foreignkey:ID;association_foreignkey:ID;" json:"author"`

	// 作者 ID
	AuthorID uint `gorm:"not null;default:0;comment:'作者 ID'" json:"author_id"`

	// 标签
	TagModel []TagModel `gorm:"many2many:article_tag;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:ID;jointable_foreignkey:ID;" json:"TagModel"`

	// 文章标签
	Tags _type.TagsArray `gorm:"type: varchar(20);not null;default:'';comment:'文章标签'" json:"tags"`

	// 来源
	Source string `gorm:"type:varchar(20);not null;default:'';comment:'来源'" json:"source"`

	// 原文链接
	SourceUrl string `gorm:"type:varchar(100);not null;default:'';comment:'原文链接'" json:"source_url"`

	// 文章封面
	Banner BannerModel `gorm:"type:varchar(100);not null;default:'';comment:'文章封面'" json:"banner"`

	// 文章封面 ID
	BannerID uint `gorm:"not null;default:0;comment:'文章封面 ID'" json:"banner_id"`

	// 文章封面 PATH
	BannerPath string `gorm:"type:varchar(100);not null;default:'';comment:'文章封面 PATH'" json:"banner_path"`
}
