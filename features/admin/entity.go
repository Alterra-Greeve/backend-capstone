package admin

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Admin struct {
	ID        string
	Name      string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminLogin struct {
	Email    string
	Password string
	Token    string
}

type AdminUpdate struct {
	ID        string
	Name      string
	Username  string
	Email     string
	Password  string
	UpdatedAt time.Time
	Token     string
}

type AdminHandlerInterface interface {
	Login() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAdminData() echo.HandlerFunc
}

type AdminServiceInterface interface {
	Login(Admin) (AdminLogin, error)
	Update(AdminUpdate) (AdminUpdate, error)
	Delete(Admin) error

	GetAdminData(Admin) (Admin, error)
}

type AdminDataInterface interface {
	Login(Admin) (Admin, error)
	Update(AdminUpdate) (Admin, error)
	Delete(Admin) error

	GetAdminByID(id string) (Admin, error)
	IsEmailExist(email string) bool
}
