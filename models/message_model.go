package models

type MessageModel struct {
	Model

	// 消息类型
	Type string `gorm:"type:varchar(20);default:''" json:"type"`

	// 消息内容
	Content string `gorm:"type:varchar(1000);default:''" json:"content"`

	// 是否已读
	IsRead bool `gorm:"not null;default:false;comment:'是否已读'" json:"is_read"`

	// 消息发送者
	Sender UserModel `gorm:"foreignkey:ID;association_foreignkey:ID;" json:"sender"`

	// 消息发送者 ID
	SenderID uint `gorm:"not null;default:0;comment:'消息发送者 ID'" json:"sender_id"`

	// 消息发送者 Name
	SenderName string `gorm:"type:varchar(20);default:''" json:"sender_name"`

	// 消息发送者 Avatar
	SenderAvatar string `gorm:"type:varchar(100);default:''" json:"sender_avatar"`

	// 消息接收者
	Receiver UserModel `gorm:"foreignkey:ID;association_foreignkey:ID;" json:"receiver"`

	// 消息接收者 ID
	ReceiverID uint `gorm:"not null;default:0;comment:'消息接收者 ID'" json:"receiver_id"`

	// 消息接收者 Name
	ReceiverName string `gorm:"type:varchar(20);default:''" json:"receiver_name"`

	// 消息接收者 Avatar
	ReceiverAvatar string `gorm:"type:varchar(100);default:''" json:"receiver_avatar"`
}
