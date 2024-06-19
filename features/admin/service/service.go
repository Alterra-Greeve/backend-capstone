package service

import (
	"backendgreeve/constant"
	"backendgreeve/features/admin"
	"backendgreeve/helper"
	"strings"
)

type AdminService struct {
	d admin.AdminDataInterface
	j helper.JWTInterface
	m helper.MailerInterface
}

func New(data admin.AdminDataInterface, jwt helper.JWTInterface, mailer helper.MailerInterface) admin.AdminServiceInterface {
	return &AdminService{
		d: data,
		j: jwt,
		m: mailer,
	}
}

func (s *AdminService) Login(userAdmin admin.Admin) (admin.AdminLogin, error) {
	if userAdmin.Email == "" || userAdmin.Password == "" {
		return admin.AdminLogin{}, constant.ErrEmptyEmailandPasswordAdmin
	}

	adminData, err := s.d.Login(userAdmin)
	if err != nil {
		return admin.AdminLogin{}, err
	}

	var AdminLogin helper.AdminJWT

	AdminLogin.ID = adminData.ID
	AdminLogin.Name = adminData.Name
	AdminLogin.Email = adminData.Email
	AdminLogin.Username = adminData.Username
	AdminLogin.Role = constant.RoleAdmin

	token, err := s.j.GenerateAdminJWT(AdminLogin)
	if err != nil {
		return admin.AdminLogin{}, err
	}

	var AdminLoginData admin.AdminLogin
	AdminLoginData.Token = token

	return AdminLoginData, nil
}

func (s *AdminService) Update(userAdmin admin.AdminUpdate) (admin.AdminUpdate, error) {
	if userAdmin.ID == "" {
		return admin.AdminUpdate{}, constant.ErrUpdateUser
	}

	if userAdmin.Name == "" && userAdmin.Username == "" && userAdmin.Email == "" && userAdmin.Password == "" {
		return admin.AdminUpdate{}, constant.ErrEditAdmin
	}

	if userAdmin.Username != "" {
		trimmedUsername := strings.TrimSpace(userAdmin.Username)
		isUsernameValid := helper.ValidateUsername(trimmedUsername)
		if !isUsernameValid {
			return admin.AdminUpdate{}, constant.ErrInvalidUsername
		}
		userAdmin.Username = trimmedUsername
	}

	if userAdmin.Password != "" {
		hashedPassword, err := helper.HashPassword(userAdmin.Password)
		if err != nil {
			return admin.AdminUpdate{}, err
		}
		userAdmin.Password = hashedPassword
	}

	if userAdmin.Name != "" {
		trimmedName := strings.TrimSpace(userAdmin.Name)
		isNameValid := helper.IsValidInput(trimmedName)
		if !isNameValid {
			return admin.AdminUpdate{}, constant.ErrFieldData
		}
		userAdmin.Name = trimmedName
	}

	if userAdmin.Email != "" {
		trimmedData := strings.TrimSpace(userAdmin.Email)
		userAdmin.Email = trimmedData
	}

	userData, err := s.d.Update(userAdmin)
	if err != nil {
		return admin.AdminUpdate{}, err
	}

	var AdminToken helper.AdminJWT
	AdminToken.ID = userData.ID
	AdminToken.Name = userData.Name
	AdminToken.Email = userData.Email
	AdminToken.Username = userData.Username
	AdminToken.Role = constant.RoleUser

	token, err := s.j.GenerateAdminJWT(AdminToken)
	if err != nil {
		return admin.AdminUpdate{}, err
	}

	userAdmin.Token = token

	return userAdmin, nil
}

func (s *AdminService) Delete(admin admin.Admin) error {
	if admin.ID == "" {
		return constant.ErrDeleteUser
	}
	return s.d.Delete(admin)
}

func (s *AdminService) GetAdminData(adminUser admin.Admin) (admin.Admin, error) {
	if adminUser.ID == "" {
		return admin.Admin{}, constant.ErrGetDataAdmin
	}
	return s.d.GetAdminByID(adminUser.ID)
}
