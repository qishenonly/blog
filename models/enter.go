package models

type Model struct {
	// 主键 ID
	ID uint `gorm:"primary_key" json:"id"`

	// 创建时间
	CreatedAt int64 `json:"created_at"`

	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}
