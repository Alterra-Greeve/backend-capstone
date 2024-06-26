package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/users"
	"backendgreeve/helper"
	"fmt"
	"math/rand"
	"strings"
)

type UserService struct {
	d users.UserDataInterface
	j helper.JWTInterface
	m helper.MailerInterface
}

func New(data users.UserDataInterface, jwt helper.JWTInterface, mailer helper.MailerInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		j: jwt,
		m: mailer,
	}
}

func (s *UserService) Register(user users.User) error {
	switch {
	case user.Email == "":
		return constant.ErrEmptyEmailRegister
	case user.Password == "":
		return constant.ErrEmptyPasswordRegister
	case user.Name == "":
		return constant.ErrEmptyNameRegister
	}
	user.Email = strings.ToLower(user.Email)
	isEmailValid := helper.ValidateEmail(user.Email)
	if !isEmailValid {
		return constant.ErrInvalidEmail
	}
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {

		return err
	}

	user.Password = hashedPassword
	user.Username = "user_" + helper.GenerateRandomString(8)

	err = s.d.Register(user)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(user users.User) (users.UserLogin, error) {
	if user.Email == "" || user.Password == "" {
		return users.UserLogin{}, constant.ErrEmptyLogin
	}
	isEmailValid := helper.ValidateEmail(user.Email)
	if !isEmailValid {
		return users.UserLogin{}, constant.ErrInvalidEmail
	}
	user.Email = strings.ToLower(user.Email)

	userData, err := s.d.Login(user)
	if err != nil {
		return users.UserLogin{}, err
	}

	var UserLogin helper.UserJWT

	UserLogin.ID = userData.ID
	UserLogin.Name = userData.Name
	UserLogin.Email = userData.Email
	UserLogin.Username = userData.Username
	UserLogin.Address = userData.Address
	UserLogin.Role = constant.RoleUser

	token, err := s.j.GenerateUserJWT(UserLogin)
	if err != nil {
		return users.UserLogin{}, err
	}

	var UserLoginData users.UserLogin
	UserLoginData.Token = token

	return UserLoginData, nil
}

func (s *UserService) Update(user users.UserUpdate) (users.UserUpdate, error) {
	if user.Email == "" && user.Username == "" && user.Password == "" && user.Address == "" && user.Name == "" && user.Gender == "" && user.Phone == "" && user.AvatarURL == "" {
		return users.UserUpdate{}, constant.ErrEmptyUpdate
	}

	if user.ID == "" {
		return users.UserUpdate{}, constant.ErrUpdateUser
	}

	if user.Email != "" {
		isEmailValid := helper.ValidateEmail(user.Email)
		if !isEmailValid {
			return users.UserUpdate{}, constant.ErrInvalidEmail
		}
		user.Email = strings.ToLower(user.Email)
	}

	if user.Username != "" {
		trimmedUsername := strings.TrimSpace(user.Username)
		isUsernameValid := helper.ValidateUsername(trimmedUsername)
		if !isUsernameValid {
			return users.UserUpdate{}, constant.ErrInvalidUsername
		}
		user.Username = trimmedUsername
	}

	if user.Phone != "" {
		trimmedPhone := strings.TrimSpace(user.Phone)
		isPhoneValid := helper.ValidatePhone(trimmedPhone)
		if !isPhoneValid {
			return users.UserUpdate{}, constant.ErrInvalidPhone
		}
		user.Phone = trimmedPhone
	}

	if user.Password != "" {
		hashedPassword, err := helper.HashPassword(user.Password)
		if err != nil {
			return users.UserUpdate{}, err
		}
		user.Password = hashedPassword
	}

	if !helper.IsValidInput(user.Name) || !helper.IsValidInput(user.Address) || !helper.IsValidInput(user.Gender) {
		return users.UserUpdate{}, constant.ErrFieldData
	}

	userData, err := s.d.Update(user)
	if err != nil {
		return users.UserUpdate{}, err
	}

	var UserToken helper.UserJWT
	UserToken.ID = userData.ID
	UserToken.Name = userData.Name
	UserToken.Email = userData.Email
	UserToken.Username = userData.Username
	UserToken.Address = userData.Address
	UserToken.Role = constant.RoleUser

	token, err := s.j.GenerateUserJWT(UserToken)
	if err != nil {
		return users.UserUpdate{}, err
	}

	user.Token = token

	return user, nil
}

func (s *UserService) GetUserData(user users.User) (users.User, error) {
	return s.d.GetUserByID(user.ID)
}

func (s *UserService) GetUserByUsername(username string) (users.User, error) {
	return s.d.GetUserByUsername(username)
}

func (s *UserService) Delete(user users.User) error {
	if user.ID == "" {
		return constant.ErrDeleteUser
	}
	return s.d.Delete(user)
}

func (s *UserService) ForgotPassword(user users.User) error {
	if user.Email == "" {
		return constant.ErrEmptyEmail
	}
	var ForgotPasswordData users.ForgotPassword
	ForgotPasswordData.Email = user.Email
	ForgotPasswordData.OTP = fmt.Sprintf("%06d", rand.Intn(1000000))

	err := s.d.ForgotPassword(ForgotPasswordData)
	if err != nil {
		return err
	}

	err = s.m.Send(user.Email, ForgotPasswordData.OTP)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) VerifyOTP(verifyOTP users.VerifyOTP) error {
	switch {
	case verifyOTP.Email == "":
		return constant.ErrEmptyEmail
	case verifyOTP.OTP == "":
		return constant.ErrEmptyOTP
	}

	err := s.d.VerifyOTP(verifyOTP)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) ResetPassword(userResetPassword users.UserResetPassword) error {
	if userResetPassword.Email == "" || userResetPassword.OTP == "" || userResetPassword.Password == "" {
		return constant.ErrEmptyResetPassword
	}

	if userResetPassword.Password != userResetPassword.ConfirmationPassword {
		return constant.ErrPasswordNotMatch
	}

	hashedPassword, err := helper.HashPassword(userResetPassword.Password)
	if err != nil {
		return err
	}
	userResetPassword.Password = hashedPassword

	return s.d.ResetPassword(userResetPassword)
}

func (s *UserService) RegisterMembership(userId string) error {
	return s.d.RegisterMembership(userId)
}

// Leaderboard
func (s *UserService) GetLeaderboard() ([]users.Leaderboard, error) {
	return s.d.GetLeaderboard()
}

// Admin
func (s *UserService) GetAllUsersForAdmin() ([]users.User, error) {
	return s.d.GetAllUsersForAdmin()
}

func (s *UserService) GetAllByPageForAdmin(page int) ([]users.User, int, error) {
	return s.d.GetAllByPageForAdmin(page)
}
func (s *UserService) GetUserByIDForAdmin(id string) (users.User, error) {
	if id == "" {
		return users.User{}, constant.ErrUserIDNotFound
	}
	return s.d.GetUserByIDForAdmin(id)
}

func (s *UserService) UpdateUserForAdmin(user users.UpdateUserByAdmin) error {
	if user.ID == "" {
		return constant.ErrUserIDNotFound
	}
	if user.Address == "" || user.Name == "" || user.Gender == "" || user.Phone == "" {
		return constant.ErrEditUserByAdmin
	}
	return s.d.UpdateUserForAdmin(user)
}

func (s *UserService) DeleteUserForAdmin(userID string) error {
	return s.d.DeleteUserForAdmin(userID)
}
func (s *UserService) GetUserImpactPointById(id string) (int, error) {
	return s.d.GetUserImpactPointById(id)
}
func (s *UserService) GetUserImpactPointByUsername(username string) (int, error) {
	return s.d.GetUserImpactPointByUsername(username)
}
