package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Username  string
	Password  string
	Address   string
	Gender    string
	Phone     string
	Coin      int
	Exp       int
	AvatarURL string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VerifyOTP struct {
	Email     string
	OTP       string
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLogin struct {
	Email    string
	Password string
	Token    string
}

type UserResetPassword struct {
	Email                string
	Password             string
	ConfirmationPassword string
	OTP                  string
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	ForgotPassword() echo.HandlerFunc
	VerifyOTP() echo.HandlerFunc
	ResetPassword() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(User) error
	Login(User) (UserLogin, error)
	Update(User) (User, error)
	Delete(User) error
	ForgotPassword(User) error
	VerifyOTP(VerifyOTP) error
	ResetPassword(UserResetPassword) error
}

type UserDataInterface interface {
	Register(User) error
	Login(User) error
	Update(User) (User, error)
	Delete(User) error
	ForgotPassword(User) error
	VerifyOTP(VerifyOTP) error
	ResetPassword(UserResetPassword) error
	IsUsernameExist(username string) bool
	IsEmailExist(email string) bool
}
