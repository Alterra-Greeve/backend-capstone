package seeds

import (
	admin "backendgreeve/features/admin/data"
	helper "backendgreeve/helper"

	"gorm.io/gorm"
)

func CreateAdminLogin(db *gorm.DB, id string, name string, username string, email string, password string) error {
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return err
	}

	adminRecord := admin.Admin{
		ID:       id,
		Name:     name,
		Password: hashedPassword,
		Username: username,
		Email:    email,
	}

	return db.Where("id = ?", id).FirstOrCreate(&adminRecord).Error
}
