package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/transaction"
	"backendgreeve/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	s transaction.TransactionServiceInterface
	j helper.JWTInterface
}

func New(s transaction.TransactionServiceInterface, j helper.JWTInterface) transaction.TransactionHandlerInterface {
	return &TransactionHandler{
		s: s,
		j: j,
	}
}

func (h *TransactionHandler) GetUserTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		return nil
	}
}

func (h *TransactionHandler) CreateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)
		userAddress := userData[constant.JWT_ADDRESS].(string)

		transactionData := transaction.CreateTransaction{
			UserID:  userId,
			Address: userAddress,
		}
		err = h.s.CreateTransaction(transactionData)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success create transaction", transactionData))
	}
}

func (h *TransactionHandler) UpdateTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		return nil
	}
}

func (h *TransactionHandler) DeleteTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		return nil
	}
}
