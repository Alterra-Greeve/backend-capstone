package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/orders"
	"backendgreeve/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrdersProductHandler struct {
	s orders.OrdersServiceInterface
	j helper.JWTInterface
}

func New(s orders.OrdersServiceInterface, j helper.JWTInterface) orders.OrdersHandlerInterface {
	return &OrdersProductHandler{
		s: s,
		j: j,
	}
}

func (h *OrdersProductHandler) GetOrdersProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := h.s.GetOrdersProduct()
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, err.Error(), nil))
		}
		if len(data) == 0 {
			helper.ObjectFormatResponse(true, constant.GetDataSuccess, nil)
		}
		response := ToResponse(data)
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.GetDataSuccess, response))
	}
}

func (h *OrdersProductHandler) GetOrdersChallenge() echo.HandlerFunc {
	return func(c echo.Context) error {
		challengeConfirmations, err := h.s.GetOrdersChallenge()
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, err.Error(), nil))
		}
		if len(challengeConfirmations) == 0 {
			helper.ObjectFormatResponse(true, constant.GetDataSuccess, nil)
		}

		response := ToChallengeConfirmationResponse(challengeConfirmations)
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.GetDataSuccess, response))
	}
}
