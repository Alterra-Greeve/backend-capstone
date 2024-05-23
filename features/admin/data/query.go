package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/admin"
	"backendgreeve/helper"
	"time"

	"gorm.io/gorm"
)

type AdminData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) admin.AdminDataInterface {
	return &AdminData{
		DB: db,
	}
}

func (u *AdminData) Login(userAdmin admin.Admin) (admin.Admin, error) {
	var data admin.Admin
	result := u.DB.Where("email = ?", userAdmin.Email).First(&data)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return admin.Admin{}, constant.ErrNotFoundEmailAdmin
		}
		return admin.Admin{}, result.Error
	}

	if !helper.CheckPasswordHash(userAdmin.Password, data.Password) {
		return admin.Admin{}, constant.ErrLoginIncorrectPassword
	}
	return data, nil
}
func (u *AdminData) Update(userAdmin admin.AdminUpdate) (admin.Admin, error) {
	var existingAdmin admin.Admin
	err := u.DB.Where("id = ?", userAdmin.ID).First(&existingAdmin).Error
	if err != nil {
		return admin.Admin{}, err
	}

	if userAdmin.Email != existingAdmin.Email {
		var count int64
		u.DB.Table("admins").Where("email = ?", userAdmin.Email).Count(&count)
		if count > 0 {
			return admin.Admin{}, err
		}
	}
	userAdmin.UpdatedAt = time.Now()

	if err := u.DB.Table("admins").Where("id = ?", userAdmin.ID).Updates(&userAdmin).Error; err != nil {
		return admin.Admin{}, err
	}

	var adminData admin.Admin
	adminData, err = u.GetAdminByID(userAdmin.ID)
	if err != nil {
		return admin.Admin{}, err
	}
	return adminData, nil
}
func (u *AdminData) Delete(userAdmin admin.Admin) error {
	_, err := u.GetAdminByID(userAdmin.ID)
	if err != nil {
		return err
	}
	if err := u.DB.Where("id = ?", userAdmin.ID).Delete(&userAdmin).Error; err != nil {
		return constant.ErrDeleteUser
	}
	return nil
}

func (u *AdminData) GetAdminDetail(email string) (*admin.Admin, error) {
	var data admin.Admin
	err := u.DB.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (u *AdminData) GetAdminByID(id string) (admin.Admin, error) {
	var userAdmin admin.Admin
	var count int64
	u.DB.Table("admins").Where("id = ?", id).Count(&count)
	if count == 0 {
		return admin.Admin{}, constant.UserNotFound
	}
	if err := u.DB.Where("id = ?", id).First(&userAdmin).Error; err != nil {
		return admin.Admin{}, constant.UserNotFound
	}
	return userAdmin, nil
}

func (u *AdminData) IsEmailExist(email string) bool {
	var userAdmin admin.Admin
	if err := u.DB.Where("email = ?", email).First(&userAdmin).Error; err != nil {
		return false
	}
	return true
}
