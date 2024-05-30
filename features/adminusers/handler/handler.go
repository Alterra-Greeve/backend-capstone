package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/adminusers"
	users "backendgreeve/features/users"
	"backendgreeve/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UsersbyAdmin struct {
	s adminusers.UserServicebyAdminInterface
	j helper.JWTInterface
}

func New(s adminusers.UserServicebyAdminInterface, j helper.JWTInterface) adminusers.UserHandlerbyAdminInterface {
	return &UsersbyAdmin{
		s: s,
		j: j,
	}
}

func (h *UsersbyAdmin) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		adminData := h.j.ExtractUserToken(token)
		role := adminData[constant.JWT_ROLE]

		if role != constant.RoleAdmin {
			helper.UnauthorizedError(c)
		}

		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		var totalPages int
		var user []users.User
		user, totalPages, err = h.s.GetAllByPage(page)

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
				Page:      totalPages,
			})
		}
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.AdminSuccessGetAllUser, response))
	}
}

func (h *UsersbyAdmin) GetUserByID() echo.HandlerFunc {
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

		users, err := h.s.GetUserByID(userId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.ObjectFormatResponse(false, err.Error(), nil))
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
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.AdminSuccessGetUser, response))
	}
}

func (h *UsersbyAdmin) UpdateUser() echo.HandlerFunc {
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
		_, err = h.s.GetUserByID(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, string(constant.ErrUserIDNotFound.Error()), nil))
		}

		var userEdit UserbyAdminRequest
		if err := c.Bind(&userEdit); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}
		response := adminusers.UpdateUserByAdmin{
			ID:      id,
			Name:    userEdit.Name,
			Address: userEdit.Address,
			Gender:  userEdit.Gender,
			Phone:   userEdit.Phone,
		}

		if err := h.s.UpdateUser(response); err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.AdminSuccessUpdateUser, nil))
	}
}

func (h *UsersbyAdmin) DeleteUser() echo.HandlerFunc {
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
		if err := h.s.DeleteUser(id); err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.AdminSuccessDeleteUser, nil))
	}
}
