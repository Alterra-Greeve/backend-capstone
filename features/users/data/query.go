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
	var UserLoginData User
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
	var userLogin users.User
	userLogin.ID = UserLoginData.ID
	userLogin.Name = UserLoginData.Name
	userLogin.Email = UserLoginData.Email
	userLogin.Username = UserLoginData.Username
	userLogin.Address = UserLoginData.Address
	return userLogin, nil
}

func (u *UserData) Update(user users.UserUpdate) (users.User, error) {
	var existingUser users.User
	err := u.DB.Where("id = ?", user.ID).First(&existingUser).Error
	if err != nil {
		return users.User{}, err
	}

	if user.Email != "" && user.Email != existingUser.Email {
		var count int64
		u.DB.Table("users").Where("email = ?", user.Email).Count(&count)
		if count > 0 {
			return users.User{}, constant.ErrEmailAlreadyExist
		}
	}

	if user.Username != "" && user.Username != existingUser.Username {
		var count int64
		u.DB.Table("users").Where("username = ?", user.Username).Count(&count)
		if count > 0 {
			return users.User{}, constant.ErrUsernameAlreadyExist
		}
	}

	if err := u.DB.Model(&User{}).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return users.User{}, constant.ErrUpdateUser
	}

	var updatedUser users.User
	updatedUser, err = u.GetUserByID(user.ID)
	if err != nil {
		return users.User{}, err
	}
	return updatedUser, nil
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
	var user User
	if err := u.DB.Model(&User{}).Where("email = ?", OTP.Email).First(&user).Error; err != nil {
		return constant.ErrEmailNotFound
	}

	OTP.ID = uuid.New().String()
	OTP.ExpiredAt = time.Now().Add(time.Minute * 5)
	verifyOTP := VerifyOTP{
		ID:        OTP.ID,
		OTP:       OTP.OTP,
		Email:     OTP.Email,
		ExpiredAt: OTP.ExpiredAt,
	}
	if err := u.DB.Create(&verifyOTP).Error; err != nil {
		return constant.ErrForgotPassword
	}

	return nil
}

func (u *UserData) VerifyOTP(verifyOTP users.VerifyOTP) error {
	var count int64
	u.DB.Model(&VerifyOTP{}).Where("email = ? AND otp = ?", verifyOTP.Email, verifyOTP.OTP).Count(&count)
	if count == 0 {
		return constant.ErrOTPNotValid
	}

	if err := u.DB.Model(&VerifyOTP{}).Where("email = ? AND otp = ? AND expired_at > ?", verifyOTP.Email, verifyOTP.OTP, time.Now()).First(&verifyOTP).Error; err != nil {
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
	var user User
	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return false
	}
	return true
}

func (u *UserData) IsEmailExist(email string) bool {
	var user User
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
	var user User

	if err := u.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return users.User{}, constant.UserNotFound
	}
	users := users.User{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Username:   user.Username,
		Address:    user.Address,
		Gender:     user.Gender,
		Coin:       user.Coin,
		Exp:        user.Exp,
		Phone:      user.Phone,
		Membership: user.Membership,
		AvatarURL:  user.AvatarURL,
	}
	return users, nil
}

func (u *UserData) RegisterMembership(userId string) error {
	var user User
	if err := u.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return constant.UserNotFound
	}
	user.Membership = true
	return u.DB.Model(&user).Update("membership", true).Error
}
func (u *UserData) GetLeaderboard() ([]users.Leaderboard, error) {
	var usersLeaderboard []users.Leaderboard
	res := u.DB.Table("users").Where("deleted_at IS NULL").Order("exp DESC").Limit(20).Scan(&usersLeaderboard)
	if res.Error != nil {
		return nil, res.Error
	}

	return usersLeaderboard, nil
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
	count := u.DB.Model(&User{}).Count(&total)
	if count.Error != nil {
		return nil, 0, constant.ErrUserDataEmpty
	}

	dataforumPerPage := 20
	totalPages := int((total + int64(dataforumPerPage) - 1) / int64(dataforumPerPage))

	tx := u.DB.Model(&User{}).Offset((page - 1) * dataforumPerPage).
		Find(&forum)
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
	res := u.DB.Model(&User{}).Where("id = ? AND deleted_at IS NULL", id).First(&users)
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

	tx := u.DB.Model(&existingUser).Omit("CreatedAt").Where("id = ?", existingUser.ID).Save(&existingUser)
	if tx.Error != nil {
		return constant.ErrUpdateUser
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

func (u *UserData) GetUserImpactPointById(userId string) (int, error) {
	var totalProductImpactPoint int
	u.DB.Table("transactions").
		Joins("JOIN transaction_items ON transactions.id = transaction_items.transaction_id").
		Joins("JOIN products ON transaction_items.product_id = products.id").
		Joins("JOIN product_impact_categories ON products.id = product_impact_categories.product_id").
		Joins("JOIN impact_categories ON product_impact_categories.impact_category_id = impact_categories.id").
		Where("transactions.user_id = ?", userId).
		Where("transactions.status = 'Berhasil'").
		Select("COALESCE(SUM(impact_categories.impact_point), 0) AS totalProductImpactPoint").
		Scan(&totalProductImpactPoint)

	var totalChallengeImpactPoint int
	u.DB.Table("challenges").
		Joins("JOIN challenge_confirmations ON challenges.id = challenge_confirmations.challenge_id").
		Joins("JOIN challenge_impact_categories ON challenges.id = challenge_impact_categories.challenge_id").
		Joins("JOIN impact_categories ON challenge_impact_categories.impact_category_id = impact_categories.id").
		Where("challenge_confirmations.status = 'Diterima' AND challenge_confirmations.user_id = ?", userId).
		Select("COALESCE(SUM(impact_categories.impact_point), 0) AS totalChallengeImpactPoint").
		Scan(&totalChallengeImpactPoint)

	impactPoint := totalProductImpactPoint + totalChallengeImpactPoint
	if impactPoint == 0 {
		return 0, nil
	}
	return impactPoint, nil
}

func (u *UserData) GetUserImpactPointByUsername(username string) (int, error) {
	// Hitung total impact point dari pembelian produk
	var totalProductImpactPoint int
	u.DB.Table("transactions").
		Joins("JOIN transaction_items ON transactions.id = transaction_items.transaction_id").
		Joins("JOIN product ON transaction_items.product_id = product.id").
		Joins("JOIN product_impact_categories ON product.id = product_impact_categories.product_id").
		Joins("JOIN impact_categories ON product_impact_categories.impact_categories_id = impact_categories.id").
		Where("transactions.user_id = (SELECT id FROM users WHERE username = ?)", username).
		Where("transaction_items.status = 'capture' OR transaction_items.status = 'accept' OR transaction_items.status = 'settlement'").
		Select("SUM(impact_categories.impact_point)").
		Scan(&totalProductImpactPoint)

	// Hitung total impact point dari challenge yang diikuti
	var totalChallengeImpactPoint int
	u.DB.Table("challenges").
		Joins("JOIN challenge_confirmation ON challenges.id = challenge_confirmation.challenge_id").
		Joins("JOIN challenge_impact_categories ON challenges.id = challenge_impact_categories.challenge_id").
		Joins("JOIN impact_categories ON challenge_impact_categories.impact_category_id = impact_categories.id").
		Where("challenge_confirmation.status = 'Diterima' AND challenge_confirmation.user_id = (SELECT id FROM users WHERE username = ?)", username).
		Select("COALESCE(SUM(impact_categories.impact_point), 0) AS totalImpactPoints").
		Scan(&totalChallengeImpactPoint)

	impactPoint := totalProductImpactPoint + totalChallengeImpactPoint
	if impactPoint == 0 {
		return 0, nil
	}
	return impactPoint, nil
}

func (u *UserData) AddReccomendationToUser(userId string, category []string) error {
	tx := u.DB.Begin()
	err := u.DB.Model(&UserReccomendation{}).Where("user_id = ?", userId).Delete(&UserReccomendation{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, cat := range category {
		newUserReccomendation := UserReccomendation{
			ID:               uuid.New().String(),
			UserID:           userId,
			ImpactCategoryID: cat,
		}
		if err := u.DB.Create(&newUserReccomendation).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
