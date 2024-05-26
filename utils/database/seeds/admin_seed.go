package seeds

import (
	admin "backendgreeve/features/admin/data"
	helper "backendgreeve/helper"

	"gorm.io/gorm"
)

func CreateAdminLogin(db *gorm.DB, id string, name string, password string, email string, username string) error {
	hashedPassword, err := helper.HashPassword(password)
	if err != nil {
		return err
	}

	adminRecord := admin.Admin{
		ID:       id,
		Name:     name,
		Password: hashedPassword,
		Email:    email,
		Username: username,
	}

	return db.Where("id = ?", id).FirstOrCreate(&adminRecord).Error
}
