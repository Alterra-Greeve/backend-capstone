package seeds

import (
	user "backendgreeve/features/users/data"
	helper "backendgreeve/helper"

	"gorm.io/gorm"
)

func CreateUsers(db *gorm.DB, user user.User) error {
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return db.Where("id = ?", user.ID).FirstOrCreate(&user).Error
}
