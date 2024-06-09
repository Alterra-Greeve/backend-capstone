package seeds

import (
	user "backendgreeve/features/users/data"
	helper "backendgreeve/helper"

	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB, id, name, password, email, username, address, gender, phone string, coin, exp int, avatar_url string) error {
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return err
	}

	usersRecord := user.User{
		ID:        id,
		Name:      name,
		Password:  hashedPassword,
		Email:     email,
		Username:  username,
		Address:   address,
		Gender:    gender,
		Phone:     phone,
		Coin:      coin,
		Exp:       exp,
		AvatarURL: avatar_url,
	}

	return db.Where("id = ?", id).FirstOrCreate(&usersRecord).Error
}
