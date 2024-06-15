package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/users"
	"backendgreeve/helper"
	"net/http"
	"strconv"
	"time"

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
			Password: UserRegisterRequest.Password,
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
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.ErrUpdateUser.Error(), nil))
		}

		currentUser, err := h.s.GetUserByIDForAdmin(userId.(string))
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		if UserUpdateRequest.Email == "" || UserUpdateRequest.Email == currentUser.Email {
			UserUpdateRequest.Email = currentUser.Email
		}
		if UserUpdateRequest.Username == "" || UserUpdateRequest.Username == currentUser.Username {
			UserUpdateRequest.Username = currentUser.Username
		}
		if UserUpdateRequest.AvatarURL == "" || UserUpdateRequest.AvatarURL == currentUser.AvatarURL {
			UserUpdateRequest.AvatarURL = currentUser.AvatarURL
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

		impactPoint, err := h.s.GetUserImpactPointById(user.ID)
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
		response.Coin = user.Coin
		response.Exp = user.Exp
		response.ImpactPoint = impactPoint
		response.Membership = user.Membership
		response.AvatarURL = user.AvatarURL
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.UserSuccessGetUser, response))
	}
}

func (h *UserHandler) GetUserByUsername() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		username := c.Param("username")

		user, err := h.s.GetUserByUsername(username)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		impactPoint, err := h.s.GetUserImpactPointByUsername(username)
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
		response.Coin = user.Coin
		response.Exp = user.Exp
		response.ImpactPoint = impactPoint
		response.Membership = user.Membership
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

func (h *UserHandler) RegisterMembership() echo.HandlerFunc {
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

		err = h.s.RegisterMembership(userId.(string))
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success", nil))
	}
}

// Leaderboard
func (h *UserHandler) GetLeaderboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		users := h.j.ExtractUserToken(token)
		role := users[constant.JWT_ROLE]

		if role != constant.RoleUser {
			return helper.UnauthorizedError(c)
		}

		voucher, err := h.s.GetLeaderboard()
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response []UserLeaderboardResponse
		for i, vouchers := range voucher {
			response = append(response, UserLeaderboardResponse{
				ID:        vouchers.ID,
				Name:      vouchers.Name,
				Username:  vouchers.Username,
				Exp:       vouchers.Exp,
				AvatarURL: vouchers.AvatarURL,
				Rank:      i + 1,
			})
		}
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.SuccessGetLeaderboard, response))
	}
}

// Admin
func (h *UserHandler) GetAllUsersForAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		adminData := h.j.ExtractAdminToken(token)
		role := adminData[constant.JWT_ROLE]

		if role != constant.RoleAdmin {
			helper.UnauthorizedError(c)
		}

		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
		var totalPages int
		var user []users.User
		user, totalPages, err = h.s.GetAllByPageForAdmin(page)

		metadata := MetadataResponse{
			CurrentPage: page,
			TotalPage:   totalPages,
		}

		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		var response []UserbyAdminandPageResponse
		for _, user := range user {
			response = append(response, UserbyAdminandPageResponse{
				ID:        user.ID,
				Name:      user.Name,
				Email:     user.Email,
				Username:  user.Username,
				Password:  user.Password,
				Address:   user.Address,
				Gender:    user.Gender,
				Phone:     user.Phone,
				Coin:      user.Coin,
				Exp:       user.Exp,
				AvatarURL: user.AvatarURL,
				CreatedAt: user.CreatedAt.Format("02/01/06"),
				UpdatedAt: user.UpdatedAt.Format("02/01/06"),
			})
		}
		return c.JSON(http.StatusOK, helper.MetadataFormatResponse(true, constant.AdminSuccessGetAllUser, metadata, response))
	}
}

func (h *UserHandler) GetUserByIDForAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		adminData := h.j.ExtractAdminToken(token)
		role := adminData[constant.JWT_ROLE]

		if role != constant.RoleAdmin {
			return helper.UnauthorizedError(c)
		}

		userId := c.Param("id")
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		users, err := h.s.GetUserByIDForAdmin(userId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(constant.ErrUserIDNotFound), helper.ObjectFormatResponse(false, constant.ErrUserIDNotFound.Error(), nil))
		}

		response := UserbyAdminResponse{
			ID:        users.ID,
			Name:      users.Name,
			Email:     users.Email,
			Username:  users.Username,
			Password:  users.Password,
			Address:   users.Address,
			Gender:    users.Gender,
			Phone:     users.Phone,
			Coin:      users.Coin,
			Exp:       users.Exp,
			AvatarURL: users.AvatarURL,
			CreatedAt: users.CreatedAt.Format("02/01/06"),
			UpdatedAt: users.UpdatedAt.Format("02/01/06"),
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.AdminSuccessGetUser, response))
	}
}

func (h *UserHandler) UpdateUserForAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		adminData := h.j.ExtractAdminToken(token)
		role := adminData[constant.JWT_ROLE]

		if role != constant.RoleAdmin {
			return helper.UnauthorizedError(c)
		}

		id := c.Param("id")
		_, err = h.s.GetUserByIDForAdmin(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, string(constant.ErrUserIDNotFound.Error()), nil))
		}

		var userEdit UserbyAdminRequest
		if err := c.Bind(&userEdit); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}
		response := users.UpdateUserByAdmin{
			ID:       id,
			Name:     userEdit.Name,
			Address:  userEdit.Address,
			Gender:   userEdit.Gender,
			Phone:    userEdit.Phone,
			UpdateAt: time.Now(),
		}

		if err := h.s.UpdateUserForAdmin(response); err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.AdminSuccessUpdateUser, nil))
	}
}

func (h *UserHandler) DeleteUserForAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, constant.Unauthorized, nil))
		}

		adminData := h.j.ExtractAdminToken(token)
		role := adminData[constant.JWT_ROLE]

		if role != constant.RoleAdmin {
			return c.JSON(http.StatusUnauthorized, helper.FormatResponse(false, constant.Unauthorized, nil))
		}

		id := c.Param("id")
		if err := h.s.DeleteUserForAdmin(id); err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.AdminSuccessDeleteUser, nil))
	}
}
