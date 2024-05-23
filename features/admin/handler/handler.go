package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/admin"
	"backendgreeve/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	s admin.AdminServiceInterface
	j helper.JWTInterface
}

func New(u admin.AdminServiceInterface, j helper.JWTInterface) admin.AdminHandlerInterface {
	return &AdminHandler{
		s: u,
		j: j,
	}
}

func (h *AdminHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var AdminLoginRequest AdminLoginRequest

		err := c.Bind(&AdminLoginRequest)
		if err != nil {
			err, message := helper.HandleEchoError(err)
			return c.JSON(err, helper.FormatResponse(false, message, nil))
		}

		admin := admin.Admin{
			Email:    AdminLoginRequest.Email,
			Password: AdminLoginRequest.Password,
		}

		adminLogin, err := h.s.Login(admin)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response AdminLoginResponse
		response.Token = adminLogin.Token
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.AdminSuccessLogin, response))
	}
}
func (h *AdminHandler) Update() echo.HandlerFunc {
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

		var AdminUpdateRequest AdminUpdateRequest
		err = c.Bind(&AdminUpdateRequest)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		admin := admin.AdminUpdate{
			ID:       userId.(string),
			Username: AdminUpdateRequest.Username,
			Name:     AdminUpdateRequest.Name,
			Email:    AdminUpdateRequest.Email,
			Password: AdminUpdateRequest.Password,
		}
		_, err = h.s.Update(admin)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.AdminSuccessUpdate, nil))
	}
}
func (h *AdminHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		admin := c.Get("admin").(admin.Admin)
		err := h.s.Delete(admin)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(true, constant.AdminSuccessDelete, nil))
	}
}

func (h *AdminHandler) GetAdminData() echo.HandlerFunc {
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
		userId := adminData[constant.JWT_ID]
		var admin admin.Admin
		admin.ID = userId.(string)

		admin, err = h.s.GetAdminData(admin)

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response AdminInfoResponse
		response.ID = admin.ID
		response.Name = admin.Name
		response.Email = admin.Email
		response.Username = admin.Username
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.UserSuccessGetUser, response))
	}
}
