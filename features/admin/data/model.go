package data

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	*gorm.Model
	ID        string          `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name      string          `gorm:"type:varchar(255);not null;column:name"`
	Username  string          `gorm:"type:varchar(255);not null;column:username;unique"`
	Email     string          `gorm:"type:varchar(255);not null;column:email;unique"`
	Password  string          `gorm:"type:varchar(255);not null;column:password"`
	CreatedAt time.Time       `gorm:"not null;column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *Admin) TableName() string {
	return "admins"
}
