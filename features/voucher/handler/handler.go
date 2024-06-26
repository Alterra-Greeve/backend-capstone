package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/voucher"
	"backendgreeve/helper"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type VoucherHandler struct {
	s voucher.VoucherServiceInterface
	j helper.JWTInterface
}

func New(s voucher.VoucherServiceInterface, j helper.JWTInterface) voucher.VoucherHandlerInterface {
	return &VoucherHandler{
		s: s,
		j: j,
	}
}

// Voucher
func (h *VoucherHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}
		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}
		var voucher []voucher.Voucher
		voucher, err = h.s.GetAll()
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response []VoucherResponse
		for i, vouchers := range voucher {
			response = append(response, VoucherResponse{
				ID:          vouchers.ID,
				Name:        vouchers.Name,
				Code:        vouchers.Code,
				Discount:    vouchers.Discount,
				Description: vouchers.Description,
				ExpiredAt:   vouchers.ExpiredAt.Format("02/01/2006 15:04"),
			})
			used, err := h.s.GetVoucherUsed(vouchers.ID)
			if err != nil {
				return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
			}
			response[i].Used = used
		}
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.VoucherSuccessGetAll, response))
	}
}

func (h *VoucherHandler) GetByIdVoucher() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		voucherId := c.Param("id")
		voucher, err := h.s.GetByIdVoucher(voucherId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(constant.ErrGetVoucherById), helper.ObjectFormatResponse(false, constant.ErrGetVoucherById.Error(), nil))
		}

		response := VoucherResponse{
			ID:          voucher.ID,
			Name:        voucher.Name,
			Code:        voucher.Code,
			Discount:    voucher.Discount,
			Description: voucher.Description,
			ExpiredAt:   voucher.ExpiredAt.Format("02/01/2006 15:04"),
		}
		used, err := h.s.GetVoucherUsed(voucher.ID)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		response.Used = used

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.VoucherSuccessGetByID, response))
	}
}

func (h *VoucherHandler) Create() echo.HandlerFunc {
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

		var voucherRequest VoucherRequest
		err = c.Bind(&voucherRequest)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		if !helper.ValidateDiscount(voucherRequest.Discount) {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Error Discount Format", nil))
		}

		if len(voucherRequest.Code) != 9 {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.ErrCodeVoucher.Error(), nil))
		}

		if len(voucherRequest.Code) == 9 && !helper.ValidateCode(voucherRequest.Code) {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Format Code Wrong", nil))
		}

		voucherData := voucher.Voucher{
			ID:          uuid.New().String(),
			Name:        voucherRequest.Name,
			Discount:    voucherRequest.Discount,
			Description: voucherRequest.Description,
			ExpiredAt:   voucherRequest.ExpiredAt,
			Code:        voucherRequest.Code,
		}
		err = h.s.Create(voucherData)
		if err != nil {
			if err == constant.ErrCodeVoucherExists {
				return c.JSON(helper.ConvertResponseCode(constant.ErrCodeVoucherExists), helper.FormatResponse(false, constant.ErrCodeVoucherExists.Error(), nil))
			}
			if err == constant.ErrVoucherField {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.ErrVoucherField.Error(), nil))
			}
			return c.JSON(helper.ConvertResponseCode(constant.ErrCreateVoucher), helper.FormatResponse(false, constant.ErrCreateVoucher.Error(), nil))
		}
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.VoucherSuccessCreate, nil))
	}
}

func (h *VoucherHandler) Update() echo.HandlerFunc {
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
		existingVoucher, err := h.s.GetByIdVoucher(id)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(constant.ErrGetVoucherById), helper.FormatResponse(false, constant.ErrGetVoucherById.Error(), nil))
		}

		var voucherRequest VoucherRequest
		err = c.Bind(&voucherRequest)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		if !helper.ValidateDiscount(voucherRequest.Discount) {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Error Discount Format", nil))
		}

		if voucherRequest.Code != "" && len(voucherRequest.Code) != 9 {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.ErrCodeVoucher.Error(), nil))
		}

		if len(voucherRequest.Code) == 9 && !helper.ValidateCode(voucherRequest.Code) {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Format Code Wrong", nil))
		}
		voucherData := voucher.VoucherEdit{
			ID:          existingVoucher.ID,
			Name:        voucherRequest.Name,
			Discount:    voucherRequest.Discount,
			Description: voucherRequest.Description,
			ExpiredAt:   voucherRequest.ExpiredAt,
			Code:        voucherRequest.Code,
		}

		err = h.s.Update(voucherData)
		if err != nil {
			if err == constant.ErrCodeVoucherExists {
				return c.JSON(helper.ConvertResponseCode(constant.ErrCodeVoucherExists), helper.FormatResponse(false, constant.ErrCodeVoucherExists.Error(), nil))
			}
			if err == constant.ErrVoucherField {
				return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, constant.ErrVoucherField.Error(), nil))
			}
			return c.JSON(helper.ConvertResponseCode(constant.ErrUpdateVoucher), helper.FormatResponse(false, constant.ErrUpdateVoucher.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.VoucherSuccessUpdate, nil))
	}
}

func (h *VoucherHandler) Delete() echo.HandlerFunc {
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

		voucherDataId := c.Param("id")
		err = h.s.Delete(voucherDataId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(constant.ErrDeleteVoucher), helper.FormatResponse(false, constant.ErrDeleteVoucher.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.VoucherSuccessDelete, nil))
	}
}
