package models

type LoginDataModel struct {
	Model

	// User
	User UserModel `gorm:"foreignkey:UserID;association_foreignkey:ID;" json:"user"`

	// User ID
	UserID uint `gorm:"not null;default:0;comment:'User ID'" json:"user_id"`

	// IP
	IP string `gorm:"type:varchar(20);default:''" json:"ip"`

	// 地址
	Address string `gorm:"type:varchar(100);default:''" json:"address"`

	// Device
	Device string `gorm:"type:varchar(100);default:''" json:"device"`

	// Token
	Token string `gorm:"type:varchar(100);default:''" json:"token"`
}
