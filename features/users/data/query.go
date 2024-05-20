package data

import (
	"backendgreeve/features/users"

	"gorm.io/gorm"
)

type UserData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		DB: db,
	}
}

func (u *UserData) Register(newUser users.User) error {
	if err := u.DB.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserData) Login(user users.User) error {
	if err := u.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserData) Update(user users.User) (users.User, error) {
	if err := u.DB.Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserData) Delete(user users.User) error {
	if err := u.DB.Where("id = ?", user.ID).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserData) ForgotPassword(user users.User) error {
	if err := u.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserData) VerifyOTP(verifyOTP users.VerifyOTP) error {
	if err := u.DB.Where("email = ?", verifyOTP.Email).First(&verifyOTP).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserData) ResetPassword(userResetPassword users.UserResetPassword) error {
	if err := u.DB.Where("email = ?", userResetPassword.Email).First(&userResetPassword).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserData) IsUsernameExist(username string) bool {
	var user users.User
	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return false
	}
	return true
}

func (u *UserData) IsEmailExist(email string) bool {
	var user users.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return false
	}
	return true
}

