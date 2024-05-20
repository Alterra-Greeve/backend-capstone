package handler

import (
	"backendgreeve/features/users"
	"backendgreeve/helper"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s users.UserServiceInterface
	j helper.JWTInterface
}

func New(u users.UserServiceInterface, j helper.JWTInterface) users.UserHandlerInterface {
	return &UserHandler{
		s: u,
		j: j,
	}
}

func (h *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserRegisterRequest UserRegisterRequest
		err := c.Bind(&UserRegisterRequest)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		user := users.User{
			Name:     UserRegisterRequest.Name,
			Email:    UserRegisterRequest.Email,
			Username: UserRegisterRequest.Username,
			Password: UserRegisterRequest.Password,
			Address:  UserRegisterRequest.Address,
			Gender:   UserRegisterRequest.Gender,
			Phone:    UserRegisterRequest.Phone,
		}
		err = h.s.Register(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, "Success", []interface{}{}))
	}
}

func (h *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserLoginRequest UserLoginRequest
		err := c.Bind(&UserLoginRequest)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		user := users.User{
			Email:    UserLoginRequest.Email,
			Password: UserLoginRequest.Password,
		}
		userLogin, err := h.s.Login(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, "Success", []interface{}{userLogin}))
	}
}

func (h *UserHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserUpdateRequest UserUpdateRequest
		err := c.Bind(&UserUpdateRequest)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		user := users.User{
			Name:     UserUpdateRequest.Name,
			Email:    UserUpdateRequest.Email,
			Username: UserUpdateRequest.Username,
			Address:  UserUpdateRequest.Address,
			Gender:   UserUpdateRequest.Gender,
			Phone:    UserUpdateRequest.Phone,
		}
		_, err = h.s.Update(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, "Success", []interface{}{}))
	}
}

func (h *UserHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(users.User)
		err := h.s.Delete(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, "Success", []interface{}{}))
	}
}

func (h *UserHandler) ForgotPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserForgotPasswordRequest UserForgotPasswordRequest
		err := c.Bind(&UserForgotPasswordRequest)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		user := users.User{
			Email: UserForgotPasswordRequest.Email,
		}
		err = h.s.ForgotPassword(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, "Success", []interface{}{}))
	}
}

func (h *UserHandler) VerifyOTP() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserVerifyOTPRequest UserVerifyOTPRequest
		err := c.Bind(&UserVerifyOTPRequest)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		user := users.VerifyOTP{
			Email: UserVerifyOTPRequest.Email,
			OTP:   UserVerifyOTPRequest.OTP,
		}
		err = h.s.VerifyOTP(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, "Success", []interface{}{}))
	}
}

func (h *UserHandler) ResetPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserResetPasswordRequest UserResetPasswordRequest
		err := c.Bind(&UserResetPasswordRequest)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		user := users.UserResetPassword{
			Email:                UserResetPasswordRequest.Email,
			Password:             UserResetPasswordRequest.Password,
			ConfirmationPassword: UserResetPasswordRequest.ConfirmationPassword,
			OTP:                  UserResetPasswordRequest.OTP,
		}
		err = h.s.ResetPassword(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), []interface{}{}))
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, "Success", []interface{}{}))
	}
}
