package data

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID        string        `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name      string        `gorm:"type:varchar(255);not null;column:name"`
	Email     string        `gorm:"type:varchar(255);not null;column:email;unique"`
	Username  string        `gorm:"type:varchar(255);not null;column:username;unique"`
	Password  string        `gorm:"type:varchar(255);not null;column:password"`
	Address   string        `gorm:"type:varchar(255);not null;column:address"`
	Gender    string        `gorm:"type:varchar(255);not null;column:gender"`
	Phone     string        `gorm:"type:varchar(255);not null;column:phone"`
	Coin      int           `gorm:"type:int;not null;column:coin"`
	Exp       int           `gorm:"type:int;not null;column:exp"`
	AvatarURL string        `gorm:"type:varchar(255);column:avatar_url;default:'https://example.com/default.jpg'"`
	CreatedAt time.Time     `gorm:"not null;column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at"`
	DeletedAt *sql.NullTime `gorm:"column:deleted_at"`
}

type VerifyOTP struct {
	*gorm.Model
	Email     string        `gorm:"type:varchar(255);not null;column:email;unique"`
	OTP       string        `gorm:"type:varchar(255);not null;column:otp"`
	ExpiredAt time.Time     `gorm:"type:varchar(255);not null;column:expired_at"`
	CreatedAt time.Time     `gorm:"not null;column:created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at"`
	DeletedAt *sql.NullTime `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *VerifyOTP) TableName() string {
	return "verify_otp"
}
