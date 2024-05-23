package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/users"
	"backendgreeve/helper"
	"log"
	"net/http"

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
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, nil))
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
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.UserSuccessRegister, nil))
	}
}

func (h *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserLoginRequest UserLoginRequest

		err := c.Bind(&UserLoginRequest)
		if err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, nil))
		}

		user := users.User{
			Email:    UserLoginRequest.Email,
			Password: UserLoginRequest.Password,
		}

		userLogin, err := h.s.Login(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response UserLoginResponse
		response.Token = userLogin.Token
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.UserSuccessLogin, response))
	}
}

func (h *UserHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID]

		var UserUpdateRequest UserUpdateRequest
		err = c.Bind(&UserUpdateRequest)
		if err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, nil))
		}

		user := users.UserUpdate{
			ID:        userId.(string),
			Name:      UserUpdateRequest.Name,
			Email:     UserUpdateRequest.Email,
			Username:  UserUpdateRequest.Username,
			Password:  UserUpdateRequest.Password,
			Address:   UserUpdateRequest.Address,
			Gender:    UserUpdateRequest.Gender,
			Phone:     UserUpdateRequest.Phone,
			AvatarURL: UserUpdateRequest.AvatarURL,
		}
		log.Println("sini")
		FromUserService, err := h.s.Update(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var UserToken UserLoginResponse
		UserToken.Token = FromUserService.Token
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.UserSuccessUpdate, UserToken))
	}
}

func (h *UserHandler) GetUserData() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID]
		var user users.User
		user.ID = userId.(string)

		user, err = h.s.GetUserData(user)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response UserInfoResponse
		response.ID = user.ID
		response.Name = user.Name
		response.Email = user.Email
		response.Username = user.Username
		response.Address = user.Address
		response.Gender = user.Gender
		response.Phone = user.Phone
		response.AvatarURL = user.AvatarURL
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.UserSuccessGetUser, response))
	}
}

func (h *UserHandler) GetUserByUsername() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		user, err := h.s.GetUserByUsername(username)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response UserInfoResponse
		response.ID = user.ID
		response.Name = user.Name
		response.Email = user.Email
		response.Username = user.Username
		response.Address = user.Address
		response.Gender = user.Gender
		response.Phone = user.Phone
		response.AvatarURL = user.AvatarURL
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.UserSuccessGetUser, response))
	}
}

func (h *UserHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID]
		var user users.User
		user.ID = userId.(string)

		err = h.s.Delete(user)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.UserSuccessDelete, nil))
	}
}

func (h *UserHandler) ForgotPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserForgotPasswordRequest UserForgotPasswordRequest

		err := c.Bind(&UserForgotPasswordRequest)

		if err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, nil))
		}

		user := users.User{
			Email: UserForgotPasswordRequest.Email,
		}
		err = h.s.ForgotPassword(user)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.UserSuccessForgotPassword, nil))
	}
}

func (h *UserHandler) VerifyOTP() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserVerifyOTPRequest UserVerifyOTPRequest
		err := c.Bind(&UserVerifyOTPRequest)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		user := users.VerifyOTP{
			Email: UserVerifyOTPRequest.Email,
			OTP:   UserVerifyOTPRequest.OTP,
		}

		err = h.s.VerifyOTP(user)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.UserSuccessOTPValidation, nil))
	}
}

func (h *UserHandler) ResetPassword() echo.HandlerFunc {
	return func(c echo.Context) error {
		var UserResetPasswordRequest UserResetPasswordRequest

		err := c.Bind(&UserResetPasswordRequest)

		if err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, nil))
		}
		user := users.UserResetPassword{
			Email:                UserResetPasswordRequest.Email,
			Password:             UserResetPasswordRequest.Password,
			ConfirmationPassword: UserResetPasswordRequest.ConfirmationPassword,
			OTP:                  UserResetPasswordRequest.OTP,
		}
		err = h.s.ResetPassword(user)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.UserSuccessResetPassword, nil))
	}
}
