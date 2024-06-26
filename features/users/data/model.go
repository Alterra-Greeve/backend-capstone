package data

import (
	impactcategory "backendgreeve/features/impactcategory/data"
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID         string `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Name       string `gorm:"type:varchar(255);column:name"`
	Email      string `gorm:"type:varchar(255);not null;column:email;unique"`
	Username   string `gorm:"type:varchar(255);column:username;unique"`
	Password   string `gorm:"type:varchar(255);not null;column:password"`
	Address    string `gorm:"type:varchar(255);column:address"`
	Gender     string `gorm:"type:varchar(255);column:gender"`
	Phone      string `gorm:"type:varchar(255);column:phone"`
	Coin       int    `gorm:"type:int;not nullcolumn:coin"`
	Exp        int    `gorm:"type:int;not null;column:exp"`
	AvatarURL  string `gorm:"type:varchar(255);column:avatar_url"`
	Membership bool   `gorm:"type:boolean;column:membership;default:false"`
}

type VerifyOTP struct {
	*gorm.Model
	ID        string    `gorm:"primary_key;type:varchar(50);not null;column:id"`
	Email     string    `gorm:"type:varchar(255);not null;column:email"`
	OTP       string    `gorm:"type:varchar(255);not null;column:otp"`
	ExpiredAt time.Time `gorm:"not null;column:expired_at"`
}

type UserReccomendation struct {
	*gorm.Model
	ID               string                        `gorm:"primary_key;type:varchar(50);not null;column:id"`
	UserID           string                        `gorm:"type:varchar(50);not null;column:user_id"`
	ImpactCategoryID string                        `gorm:"type:varchar(50);not null;column:impact_category_id"`
	ImpactCategory   impactcategory.ImpactCategory `gorm:"foreignKey:ImpactCategoryID;references:ID"`
	User             User                          `gorm:"foreignKey:UserID;references:ID"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *UserReccomendation) TableName() string {
	return "user_reccomendations"
}

func (u *VerifyOTP) TableName() string {
	return "verify_otp"
}
