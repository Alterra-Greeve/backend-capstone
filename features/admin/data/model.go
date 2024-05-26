package data

import (
	"gorm.io/gorm"
)

type Admin struct {
	*gorm.Model
	ID       string `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name     string `gorm:"type:varchar(255);not null;column:name"`
	Username string `gorm:"type:varchar(255);not null;column:username;unique"`
	Email    string `gorm:"type:varchar(255);not null;column:email;unique"`
	Password string `gorm:"type:varchar(255);not null;column:password"`
}

func (u *Admin) TableName() string {
	return "admins"
}
