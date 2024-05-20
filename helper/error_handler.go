package helper

import (
	"net/http"
	"backendgreeve/constant"

	"github.com/labstack/echo/v4"
)

func ConvertResponseCode(err error) int {
	switch err {

	// Default
	default:
		return http.StatusInternalServerError
	}
}

func HandleEchoError(err error) (int, string) {
	switch e := err.(type) {
	case *echo.HTTPError:
		return ConvertResponseCode(e), e.Message.(string)
	default:
		return ConvertResponseCode(err), err.Error()
	}
}

func UnauthorizedError(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, FormatResponse(false, constant.Unauthorized, []interface{}{}))
}
func InternalServerError(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, FormatResponse(false, constant.InternalServerError, []interface{}{}))
}
