package models

type TagModel struct {
	Model

	// 标签名称
	Name string `gorm:"type:varchar(20);default:''" json:"name"`

	// Articles
	Articles []ArticleModel `gorm:"many2many:article_tag;foreignkey:ID;association_foreignkey:ID;association_jointable_foreignkey:ID;jointable_foreignkey:ID;" json:"articles"`
}
