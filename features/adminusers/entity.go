package adminusers

import (
	user "backendgreeve/features/users"

	"github.com/labstack/echo/v4"
)

type UpdateUserByAdmin struct {
	ID      string
	Name    string
	Address string
	Gender  string
	Phone   string
}

type UserHandlerbyAdminInterface interface {
	GetAllUsers() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type UserServicebyAdminInterface interface {
	GetAllUsers() ([]user.User, error)
	GetAllByPage(page int) ([]user.User, int, error)
	GetUserByID(id string) (user.User, error)
	UpdateUser(user UpdateUserByAdmin) error
	DeleteUser(userID string) error
}

type UserDatabyAdminInterface interface {
	GetAllUsers() ([]user.User, error)
	GetAllByPage(page int) ([]user.User, int, error)
	GetUserByID(id string) (user.User, error)
	UpdateUser(user UpdateUserByAdmin) error
	DeleteUser(userID string) error
}
