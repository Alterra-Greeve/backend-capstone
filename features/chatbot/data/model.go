package data

import (
	"gorm.io/gorm"
)

type Chatbot struct {
	*gorm.Model
	ID      string `gorm:"column:id;primary_key"`
	ChatID  string `gorm:"column:chat_id;not null"`
	Role    string `gorm:"column:role;not null"`
	Message string `gorm:"column:message;not null"`
}
