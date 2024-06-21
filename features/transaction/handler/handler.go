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
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)
		transactions, err := h.s.GetUserTransaction(userId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		response := []TransactionUserResponse{}
		for _, transaction := range transactions {
			response = append(response, new(TransactionUserResponse).FromEntity(transaction))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success get user transaction", response))
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
		var request TransactionRequest
		if err := c.Bind(&request); err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		transactionData := transaction.CreateTransaction{
			UserID:      userId,
			VoucherCode: request.VoucherCode,
			UsingCoin:   request.UsingCoin,
		}
		transaction, err := h.s.CreateTransaction(transactionData)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		transactionResponse := TransactionResponse{
			ID:      transaction.ID,
			Amount:  int(transaction.Total),
			SnapURL: transaction.SnapURL,
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success create transaction", transactionResponse))
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

func (h *TransactionHandler) GetAllTransaction() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		adminData := h.j.ExtractAdminToken(token)
		role := adminData[constant.JWT_ROLE]

		if role != constant.RoleAdmin {
			return helper.UnauthorizedError(c)
		}

		transactions, err := h.s.GetAllTransaction()
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		response := []TransactionAllUserResponses{}
		for _, transaction := range transactions {
			response = append(response, new(TransactionAllUserResponses).FromEntity(transaction))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.TransactionSuccessGet, response))
	}
}
