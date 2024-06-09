package data

import (
	"backendgreeve/constant"
	"backendgreeve/features/users"
	"backendgreeve/helper"

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
	newUsers := User{
		ID:        uuid.New().String(),
		Name:      newUser.Name,
		Email:     newUser.Email,
		Username:  newUser.Username,
		Password:  newUser.Password,
		Coin:      0,
		Exp:       0,
		AvatarURL: "https://storage.googleapis.com/alterra-greeve/greeve/8aec5e90-b197-4e38-9f52-72b328259384user.png",
	}
	if err := u.DB.Create(&newUsers).Error; err != nil {
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

	if user.Email != existingUser.Email && user.Username != existingUser.Username {
		var count int64
		u.DB.Table("users").Where("email = ? OR username = ?", user.Email, user.Username).Count(&count)
		if count > 0 {
			return users.User{}, constant.ErrEmailUsernameAlreadyExist
		}
	}

	user.UpdatedAt = time.Now()
	if err := u.DB.Table("users").Where("id = ?", user.ID).Updates(&user).Error; err != nil {
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
	impactPoint, err := u.GetUserImpactPoint(user.ID)
	if err != nil {
		return users.User{}, err
	}
	user.ImpactPoint = impactPoint
	return user, nil
}

func (u *UserData) GetUserByUsername(username string) (users.User, error) {
	var user users.User

	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return users.User{}, constant.UserNotFound
	}

	return user, nil
}

// Admin
func (u *UserData) GetAllUsersForAdmin() ([]users.User, error) {
	var users []users.User
	res := u.DB.Find(&users).Where("deleted_at IS NULL")
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (u *UserData) GetAllByPageForAdmin(page int) ([]users.User, int, error) {
	var forum []users.User

	var total int64
	count := u.DB.Model(&users.User{}).Where("deleted_at IS NULL").Count(&total)
	if count.Error != nil {
		return nil, 0, constant.ErrUserDataEmpty
	}

	dataforumPerPage := 20
	totalPages := int((total + int64(dataforumPerPage) - 1) / int64(dataforumPerPage))

	tx := u.DB.Model(&User{}).Offset((page - 1) * dataforumPerPage).Limit(dataforumPerPage).Find(&forum)
	if tx.Error != nil {
		return nil, 0, constant.ErrGetUser
	}
	if tx.RowsAffected == 0 {
		return nil, 0, constant.ErrUserNotFound
	}
	return forum, totalPages, nil
}

func (u *UserData) GetUserByIDForAdmin(id string) (users.User, error) {
	var users users.User
	res := u.DB.Model(&User{}).Where(" id = ? AND deleted_at IS NULL", id).First(&users)
	if res.Error != nil {
		return users, res.Error
	}
	return users, nil
}

func (u *UserData) UpdateUserForAdmin(user users.UpdateUserByAdmin) error {
	var existingUser users.User
	if err := u.DB.Where("id = ?", user.ID).First(&existingUser).Error; err != nil {
		return err
	}

	err := u.DB.Model(&existingUser).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserData) DeleteUserForAdmin(userID string) error {
	res := u.DB.Begin()

	if err := res.Where("id = ?", userID).Delete(&User{}); err.Error != nil {
		res.Rollback()
		return constant.ErrUserIDNotFound
	} else if err.RowsAffected == 0 {
		res.Rollback()
		return constant.ErrUserIDNotFound
	}

	return res.Commit().Error
}

func (u *UserData) GetUserImpactPoint(userId string) (int, error) {
	// Hitung total impact point dari pembelian produk
	var totalProductImpactPoint int
	u.DB.Table("transactions").
		Joins("JOIN transaction_items ON transactions.id = transaction_items.transaction_id").
		Joins("JOIN product ON transaction_items.product_id = product.id").
		Joins("JOIN product_impact_categories ON product.id = product_impact_categories.product_id").
		Joins("JOIN impact_categories ON product_impact_categories.impact_categories_id = impact_categories.id").
		Where("transactions.user_id = ?", userId).
		Where("transaction_items.status = 'capture' OR transaction_items.status = 'accept' OR transaction_items.status = 'settlement'").
		Select("SUM(impact_categories.impact_point)").
		Scan(&totalProductImpactPoint)

	// Hitung total impact point dari challenge yang diikuti
	var totalChallengeImpactPoint int
	u.DB.Table("challenges").
		Joins("JOIN challenge_confirmation ON challenges.id = challenge_confirmation.challenge_id").
		Joins("JOIN challenge_impact_categories ON challenges.id = challenge_impact_categories.challenge_id").
		Joins("JOIN impact_categories ON challenge_impact_categories.impact_category_id = impact_categories.id").
		Where("challenge_confirmation.status = 'Diterima' AND challenge_confirmation.user_id = ?", userId).
		Select("COALESCE(SUM(impact_categories.impact_point), 0) AS totalImpactPoints").
		Scan(&totalChallengeImpactPoint)

	impactPoint := totalProductImpactPoint + totalChallengeImpactPoint
	if impactPoint == 0 {
		return 0, nil
	}
	return impactPoint, nil
}
