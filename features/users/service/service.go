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

func (s *UserService) Register(newUser users.User) error {
	return s.d.Register(newUser)
}

func (s *UserService) Login(user users.User) (users.UserLogin, error) {
	var UserLogin helper.UserJWT
	UserLogin.ID = user.ID
	UserLogin.Name = user.Name
	UserLogin.Email = user.Email
	UserLogin.Username = user.Username
	UserLogin.Address = user.Address
	UserLogin.Role = constant.RoleUser
	token, err := s.j.GenerateUserJWT(UserLogin)

	var UserLoginData users.UserLogin
	UserLoginData.Token = token
	return UserLoginData, err
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

