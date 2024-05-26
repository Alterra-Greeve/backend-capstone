package data

import (
	"time"

	"gorm.io/gorm"
)

type Forum struct {
	*gorm.Model
	ID          string `gorm:"primary_key;type:varchar(50);not null;column:id;"`
	Title       string `gorm:"type:varchar(255);not null;column:title"`
	AuthorID    string `gorm:"type:varchar(50);not null;column:author_id"`
	Author      User   `gorm:"foreignKey:AuthorID"`
	Description string `gorm:"type:varchar(255);not null;column:description"`
}

type User struct {
	ID        string          `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Username  string          `gorm:"type:varchar(255);not null;column:username"`
	Forums    []Forum         `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at"`
}

type MessageForum struct {
	*gorm.Model
	ID      string `gorm:"primary_key;type:varchar(50);not null;column:id"`
	UserID  string `gorm:"type:varchar(50);not null;column:user_id"`
	ForumID string `gorm:"type:varchar(50);not null;column:forum_id"`
	Message string `gorm:"type:varchar(255);not null;column:message"`
}

func (Forum) TableName() string {
	return "forums"
}

func (MessageForum) TableName() string {
	return "message_forums"
}
