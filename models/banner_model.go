package models

type BannerModel struct {
	Model

	// 路径
	Path string `json:"path"`

	// 图片Hash, 用于去重
	Hash string `gorm:"type:varchar(20);default:''" json:"hash"`

	// 名称
	Name string `gorm:"type:varchar(20);default:''" json:"name"`
}
