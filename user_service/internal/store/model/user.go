package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint32 `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"uniqueIndex;size:64;not null"`
	Password  string `gorm:"size:128;not null"`
	Phone     string `gorm:"size:128;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user_db"
}

// Message模型（shared/model/message.go）
type Message struct {
	gorm.Model
	Content    string `gorm:"type:text;not null"`
	SenderID   uint   `gorm:"index:idx_sender"`
	ReceiverID uint   `gorm:"index:idx_receiver"`
}
