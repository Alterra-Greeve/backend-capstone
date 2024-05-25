package data

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID        string `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name      string `gorm:"type:varchar(255);not null;column:name"`
	Email     string `gorm:"type:varchar(255);not null;column:email;unique"`
	Username  string `gorm:"type:varchar(255);not null;column:username;unique"`
	Password  string `gorm:"type:varchar(255);not null;column:password"`
	Address   string `gorm:"type:varchar(255);not null;column:address"`
	Gender    string `gorm:"type:varchar(255);not null;column:gender"`
	Phone     string `gorm:"type:varchar(255);not null;column:phone"`
	Coin      int    `gorm:"type:int;not null;column:coin"`
	Exp       int    `gorm:"type:int;not null;column:exp"`
	AvatarURL string `gorm:"type:varchar(255);column:avatar_url"`
}

type VerifyOTP struct {
	*gorm.Model
	ID        string    `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Email     string    `gorm:"type:varchar(255);not null;column:email"`
	OTP       string    `gorm:"type:varchar(255);not null;column:otp"`
	ExpiredAt time.Time `gorm:"not null;column:expired_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *VerifyOTP) TableName() string {
	return "verify_otp"
}
