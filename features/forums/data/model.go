package data

import (
	user "backendgreeve/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	*gorm.Model
	ID            string    `gorm:"primary_key;type:varchar(50);not null;column:id;"`
	Title         string    `gorm:"type:varchar(255);not null;column:title"`
	UserID        string    `gorm:"type:varchar(50);not null;column:user_id"`
	User          user.User `gorm:"foreignKey:UserID;references:ID"`
	Description   string    `gorm:"type:varchar(255);not null;column:description"`
	LastMessageAt time.Time `gorm:"column:last_message_at"`
	CreatedAt     time.Time `gorm:"column:created_at"`
}

type MessageForum struct {
	*gorm.Model
	ID      string    `gorm:"primary_key;type:varchar(50);not null;column:id"`
	ForumID string    `gorm:"type:varchar(50);not null;column:forum_id"`
	Forum   Forum     `gorm:"foreignKey:ForumID;references:ID"`
	Message string    `gorm:"type:varchar(255);not null;column:message"`
	UserID  string    `gorm:"type:varchar(50);not null;column:user_id"`
	User    user.User `gorm:"foreignKey:UserID;references:ID"`
}

func (Forum) TableName() string {
	return "forums"
}

func (MessageForum) TableName() string {
	return "message_forums"
}
