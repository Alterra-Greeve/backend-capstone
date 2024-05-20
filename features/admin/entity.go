package admin

import "github.com/labstack/echo/v4"

type Admin struct {
	ID       string
	Name     string
	Username string
	Email    string
	Password string
}

type AdminLogin struct {
	Email    string
	Password string
	Token    string
}

type AdminHandlerInterface interface {
	Register(Admin) echo.HandlerFunc
	Login(Admin) echo.HandlerFunc
	Update(Admin) echo.HandlerFunc
	Delete(Admin) echo.HandlerFunc
}

type AdminServiceInterface interface {
	Register(Admin) error
	Login(Admin) (AdminLogin, error)
	Update(Admin) (Admin, error)
	Delete(Admin) error
}

type AdminDataInterface interface {
	Register(Admin) error
	Login(Admin) error
	Update(Admin) (Admin, error)
	Delete(Admin) error
}
