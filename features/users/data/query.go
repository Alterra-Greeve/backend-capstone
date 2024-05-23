package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/users"
	"backendgreeve/helper"
	"log"

	"time"

	"github.com/google/uuid"
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
	isEmailExist := u.IsEmailExist(newUser.Email)
	if isEmailExist {
		return constant.ErrEmailAlreadyExist
	}

	isUsernameExist := u.IsUsernameExist(newUser.Username)
	if isUsernameExist {
		return constant.ErrUsernameAlreadyExist
	}

	if newUser.AvatarURL == "" {
		newUser.AvatarURL = "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png"
	}

	newUser.ID = uuid.New().String()
	newUser.Coin = 0
	newUser.Exp = 0
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	if err := u.DB.Create(&newUser).Error; err != nil {
		return constant.ErrRegister
	}
	return nil
}

func (u *UserData) Login(user users.User) (users.User, error) {
	var UserLoginData users.User
	result := u.DB.Where("email = ?", user.Email).First(&UserLoginData)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return users.User{}, constant.UserNotFound
		}
		return users.User{}, result.Error
	}

	if !helper.CheckPasswordHash(user.Password, UserLoginData.Password) {
		return users.User{}, constant.ErrLoginIncorrectPassword
	}

	return UserLoginData, nil
}

func (u *UserData) Update(user users.UserUpdate) (users.User, error) {
	var existingUser users.User
	err := u.DB.Where("id = ?", user.ID).First(&existingUser).Error
	if err != nil {
		return users.User{}, err
	}

	if user.Email != existingUser.Email || user.Username != existingUser.Username {
		var count int64
		u.DB.Table("users").Where("email = ? OR username = ?", user.Email, user.Username).Count(&count)
		if count > 0 {
			return users.User{}, constant.ErrEmailUsernameAlreadyExist
		}
	}

	user.UpdatedAt = time.Now()
	if err := u.DB.Table("users").Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		log.Printf("Error updating user with ID %s: %v", user.ID, err)
		log.Println(err)
		return users.User{}, constant.ErrUpdateUser
	}

	var userData users.User
	userData, err = u.GetUserByID(user.ID)
	if err != nil {
		return users.User{}, err
	}
	return userData, nil
}

func (u *UserData) Delete(user users.User) error {
	_, err := u.GetUserByID(user.ID)
	if err != nil {
		return err
	}
	if err := u.DB.Where("id = ?", user.ID).Delete(&user).Error; err != nil {
		return constant.ErrDeleteUser
	}
	return nil
}

func (u *UserData) ForgotPassword(OTP users.ForgotPassword) error {
	if err := u.DB.Table("users").Where("email = ?", OTP.Email).First(&OTP).Error; err != nil {
		return constant.ErrEmailNotFound
	}
	OTP.ID = uuid.New().String()
	OTP.ExpiredAt = time.Now().Add(time.Minute * 5)
	OTP.UpdatedAt = time.Now()

	if err := u.DB.Table("verify_otp").Create(&OTP).Error; err != nil {
		return constant.ErrForgotPassword
	}

	return nil
}

func (u *UserData) VerifyOTP(verifyOTP users.VerifyOTP) error {
	var count int64
	u.DB.Table("verify_otp").Where("email = ? AND otp = ?", verifyOTP.Email, verifyOTP.OTP).Count(&count)
	if count == 0 {
		return constant.ErrOTPNotValid
	}

	if err := u.DB.Table("verify_otp").Where("email = ? AND otp = ? AND expired_at > ?", verifyOTP.Email, verifyOTP.OTP, time.Now()).First(&verifyOTP).Error; err != nil {
		return constant.ErrOTPExpired
	}

	return nil
}

func (u *UserData) ResetPassword(userResetPassword users.UserResetPassword) error {
	var OTP users.VerifyOTP
	OTP.OTP = userResetPassword.OTP
	OTP.Email = userResetPassword.Email
	err := u.VerifyOTP(OTP)

	if err != nil {
		return err
	}

	if err := u.DB.Table("users").Where("email = ?", userResetPassword.Email).Update("password", userResetPassword.Password).Error; err != nil {
		return constant.ErrResetPassword
	}

	if err := u.DB.Table("verify_otp").Where("email = ?", userResetPassword.Email).Delete(&OTP).Error; err != nil {
		return constant.ErrResetPassword
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

func (u *UserData) GetUserByID(id string) (users.User, error) {
	var user users.User
	var count int64
	u.DB.Table("users").Where("id = ?", id).Count(&count)
	if count == 0 {
		return users.User{}, constant.UserNotFound
	}
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return users.User{}, constant.UserNotFound
	}
	return user, nil
}

func (u *UserData) GetUserByUsername(username string) (users.User, error) {
	var user users.User

	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return users.User{}, constant.UserNotFound
	}

	return user, nil
}
