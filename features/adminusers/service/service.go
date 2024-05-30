package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/adminusers"
	"backendgreeve/features/users"
	"backendgreeve/helper"
)

// V2
type UsersbyAdminService struct {
	d adminusers.UserDatabyAdminInterface
}

func New(data adminusers.UserDatabyAdminInterface, jwt helper.JWTInterface, mailer helper.MailerInterface) adminusers.UserServicebyAdminInterface {
	return &UsersbyAdminService{
		d: data,
	}
}

func (s *UsersbyAdminService) GetAllUsers() ([]users.User, error) {
	return s.d.GetAllUsers()
}

func (s *UsersbyAdminService) GetAllByPage(page int) ([]users.User, int, error) {
	return s.d.GetAllByPage(page)
}
func (s *UsersbyAdminService) GetUserByID(id string) (users.User, error) {
	if id == "" {
		return users.User{}, constant.ErrUserIDNotFound
	}
	return s.d.GetUserByID(id)
}

func (s *UsersbyAdminService) UpdateUser(user adminusers.UpdateUserByAdmin) error {
	if user.ID == "" {
		return constant.ErrUserIDNotFound
	}
	if user.Address == "" || user.Name == "" || user.Gender == "" || user.Phone == "" {
		return constant.ErrEditUserByAdmin
	}
	return s.d.UpdateUser(user)
}

func (s *UsersbyAdminService) DeleteUser(userID string) error {
	return s.d.DeleteUser(userID)
}
