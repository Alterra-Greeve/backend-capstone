package helper

import (
	"backendgreeve/constant"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ConvertResponseCode(err error) int {
	switch err {
	// General Error
	case constant.BadRequest:
		return http.StatusBadRequest

	// Users Error
	case constant.UserNotFound:
		return http.StatusNotFound
	case constant.ErrEmptyLogin:
		return http.StatusBadRequest
	case constant.ErrLoginIncorrectPassword:
		return http.StatusUnauthorized
	case constant.ErrEmptyEmailRegister:
		return http.StatusBadRequest
	case constant.ErrEmptyPasswordRegister:
		return http.StatusBadRequest
	case constant.ErrEmptyAddressRegister:
		return http.StatusBadRequest
	case constant.ErrEmptyNameRegister:
		return http.StatusBadRequest
	case constant.ErrEmptyGenderRegister:
		return http.StatusBadRequest
	case constant.ErrEmailAlreadyExist:
		return http.StatusConflict
	case constant.ErrUsernameAlreadyExist:
		return http.StatusConflict
	case constant.ErrRegister:
		return http.StatusInternalServerError
	case constant.ErrGenerateJWT:
		return http.StatusInternalServerError
	case constant.ErrValidateJWT:
		return http.StatusUnauthorized
	case constant.ErrHashPassword:
		return http.StatusInternalServerError
	case constant.ErrEmptyPhoneRegister:
		return http.StatusBadRequest
	case constant.ErrRegister:
		return http.StatusInternalServerError
	case constant.ErrUpdateUser:
		return http.StatusInternalServerError
	case constant.ErrEmptyUpdate:
		return http.StatusBadRequest
	case constant.ErrEmailUsernameAlreadyExist:
		return http.StatusConflict
	case constant.ErrEmptyEmail:
		return http.StatusBadRequest
	case constant.ErrEmailNotFound:
		return http.StatusNotFound
	case constant.ErrForgotPassword:
		return http.StatusInternalServerError
	case constant.ErrOTPNotValid:
		return http.StatusBadRequest
	case constant.ErrOTPExpired:
		return http.StatusBadRequest
	case constant.ErrEmptyOTP:
		return http.StatusBadRequest
	case constant.ErrResetPassword:
		return http.StatusInternalServerError
	case constant.ErrDeleteUser:
		return http.StatusInternalServerError
	case constant.ErrEmptyResetPassword:
		return http.StatusBadRequest
	// User Admin
	case constant.ErrUserDataEmpty:
		return http.StatusNotFound
	case constant.ErrGetUser:
		return http.StatusNotFound
	case constant.ErrUserNotFound:
		return http.StatusNotFound
	case constant.ErrUserIDNotFound:
		return http.StatusNotFound
	case constant.ErrEditUserByAdmin:
		return http.StatusBadRequest

	// Forum
	case constant.ErrGetForum:
		return http.StatusNotFound
	case constant.ErrCreateForum:
		return http.StatusBadRequest
	case constant.ErrGetForumByID:
		return http.StatusNotFound
	case constant.ErrEditForum:
		return http.StatusBadRequest
	case constant.UnatuhorizeForumAndMessage:
		return http.StatusUnauthorized
	case constant.ErrGetMessage:
		return http.StatusNotFound
	case constant.ErrGetMessageByID:
		return http.StatusBadRequest
	case constant.ErrCreateMessage:
		return http.StatusBadRequest
	case constant.ErrUpdateMessage:
		return http.StatusBadRequest
	case constant.ErrDeleteMessage:
		return http.StatusBadRequest
	case constant.ErrEditMessage:
		return http.StatusBadRequest
	case constant.ErrMessgaeNotFound:
		return http.StatusNotFound
	case constant.ErrDeleteForum:
		return http.StatusBadRequest
	case constant.ErrForumNotFound:
		return http.StatusNotFound
	case constant.ErrForumDataNotFound:
		return http.StatusNotFound

	// Admin Error
	case constant.ErrGetDataAdmin:
		return http.StatusNotFound
	case constant.ErrEmptyEmailandPasswordAdmin:
		return http.StatusBadRequest
	case constant.ErrNotFoundEmailAdmin:
		return http.StatusNotFound
	case constant.ErrEditAdmin:
		return http.StatusBadRequest
	case constant.ErrAdminEmailUsernameAlreadyExist:
		return http.StatusConflict
	case constant.ErrEmptyEmailandPasswordAdmin:
		return http.StatusBadRequest

	// Impact Category Error
	case constant.ErrImpactCategoryNotFound:
		return http.StatusNotFound
	case constant.ErrCreateImpactCategory:
		return http.StatusInternalServerError
	case constant.ErrUpdateImpactCategory:
		return http.StatusInternalServerError
	case constant.ErrDeleteImpactCategory:
		return http.StatusInternalServerError
	case constant.ErrImpactCategoryField:
		return http.StatusBadRequest
	case constant.ErrImpactCategoryFieldUpdate:
		return http.StatusBadRequest

	// Product Error
	case constant.ErrProductNotFound:
		return http.StatusNotFound
	case constant.ErrCreateProduct:
		return http.StatusInternalServerError
	case constant.ErrUpdateProduct:
		return http.StatusInternalServerError
	case constant.ErrDeleteProduct:
		return http.StatusInternalServerError
	case constant.ErrProductField:
		return http.StatusBadRequest
	case constant.ErrProductFieldUpdate:
		return http.StatusBadRequest
	case constant.ErrProductNameEmpty:
		return http.StatusBadRequest
	case constant.ErrProductDescriptionEmpty:
		return http.StatusBadRequest
	case constant.ErrProductCoinEmpty:
		return http.StatusBadRequest
	case constant.ErrProductPriceEmpty:
		return http.StatusBadRequest
	case constant.ErrProductImagesEmpty:
		return http.StatusBadRequest
	case constant.ErrProductImpactCategoriesEmpty:
		return http.StatusBadRequest
	case constant.ErrProductIDEmpty:
		return http.StatusBadRequest
	case constant.ErrProductUpdateEmpty:
		return http.StatusBadRequest
	case constant.ErrProductDelete:
		return http.StatusInternalServerError
	case constant.ErrGetProduct:
		return http.StatusInternalServerError
	case constant.ErrProductEmpty:
		return http.StatusNotFound
	// Default
	default:
		return http.StatusInternalServerError
	}
}
func HandleEchoError(err error) (int, string) {
	if _, ok := err.(*echo.HTTPError); ok {
		return http.StatusBadRequest, constant.BadInput
	}
	return http.StatusBadRequest, constant.BadInput
}

func UnauthorizedError(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, FormatResponse(false, constant.Unauthorized, nil))
}
func InternalServerError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, FormatResponse(false, constant.InternalServerError, nil))
}
func JWTErrorHandler(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, FormatResponse(false, constant.Unauthorized, nil))
}
