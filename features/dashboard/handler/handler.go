package handler

import (
	"backendgreeve/features/dashboard"
	"backendgreeve/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DashboardHandler struct {
	s dashboard.DashboardServiceInterface
	j helper.JWTInterface
}

func New(s dashboard.DashboardServiceInterface, j helper.JWTInterface) dashboard.DashboardHandlerInterface {
	return &DashboardHandler{
		s: s,
		j: j,
	}
}

func (h *DashboardHandler) GetDashboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		data, err := h.s.GetDashboard()
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		response := new(DashboardResponse).ToResponse(data)

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success", response))
	}
}
