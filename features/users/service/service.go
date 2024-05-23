package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/users"
	"backendgreeve/helper"
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
	case user.Address == "":
		return constant.ErrEmptyAddressRegister
	case user.Name == "":
		return constant.ErrEmptyNameRegister
	case user.Gender == "":
		return constant.ErrEmptyGenderRegister
	case user.Phone == "":
		return constant.ErrEmptyPhoneRegister
	}

	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

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

	err := s.d.Login(user)
	if err != nil {
		return users.UserLogin{}, err
	}

	var UserLogin helper.UserJWT

	UserLogin.ID = user.ID
	UserLogin.Name = user.Name
	UserLogin.Email = user.Email
	UserLogin.Username = user.Username
	UserLogin.Address = user.Address
	UserLogin.Role = constant.RoleUser

	token, err := s.j.GenerateUserJWT(UserLogin)
	if err != nil {
		return users.UserLogin{}, err
	}

	var UserLoginData users.UserLogin
	UserLoginData.Token = token

	return UserLoginData, nil
}

func (s *UserService) Update(user users.User) (users.User, error) {
	return s.d.Update(user)
}

func (s *UserService) Delete(user users.User) error {
	return s.d.Delete(user)
}

func (s *UserService) ForgotPassword(user users.User) error {
	return s.d.ForgotPassword(user)
}

func (s *UserService) VerifyOTP(verifyOTP users.VerifyOTP) error {
	return s.d.VerifyOTP(verifyOTP)
}

func (s *UserService) ResetPassword(userResetPassword users.UserResetPassword) error {
	return s.d.ResetPassword(userResetPassword)
}
