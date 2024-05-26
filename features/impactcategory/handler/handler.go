package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/impactcategory"
	"backendgreeve/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ImpactCategoryHandler struct {
	i impactcategory.ImpactCategoryServiceInterface
	j helper.JWTInterface
}

func New(i impactcategory.ImpactCategoryServiceInterface, j helper.JWTInterface) impactcategory.ImpactCategoryHandlerInterface {
	return &ImpactCategoryHandler{
		i: i,
		j: j,
	}
}

func (h *ImpactCategoryHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
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

		impactCategories, err := h.i.GetAll()

		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var impactCategoryResponses []ImpactCategoryResponse
		for _, impactCategory := range impactCategories {
			impactCategoryResponses = append(impactCategoryResponses, ImpactCategoryResponse{
				ID:          impactCategory.ID,
				Name:        impactCategory.Name,
				ImpactPoint: impactCategory.ImpactPoint,
				IconURL:     impactCategory.IconURL,
			})
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.ImpactCategorySuccessGet, []interface{}{impactCategoryResponses}))
	}
}

func (h *ImpactCategoryHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
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

		impactId := c.Param("id")
		impactCategory, err := h.i.GetByID(impactId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.ObjectFormatResponse(false, err.Error(), nil))
		}

		impactCategoryResponse := ImpactCategoryResponse{
			ID:          impactCategory.ID,
			Name:        impactCategory.Name,
			ImpactPoint: impactCategory.ImpactPoint,
			IconURL:     impactCategory.IconURL,
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ImpactCategorySuccessGet, impactCategoryResponse))
	}
}

func (h *ImpactCategoryHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
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

		var impactRequest CreateImpactCategoryRequest
		err = c.Bind(&impactRequest)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		impact := impactcategory.ImpactCategory{
			Name:        impactRequest.Name,
			ImpactPoint: impactRequest.ImpactPoint,
			IconURL:     impactRequest.IconURL,
		}

		err = h.i.Create(impact)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.ObjectFormatResponse(true, constant.ImpactCategorySuccessCreate, nil))
	}
}

func (h *ImpactCategoryHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
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

		impactId := c.Param("id")
		var impactRequest CreateImpactCategoryRequest
		err = c.Bind(&impactRequest)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		impact := impactcategory.ImpactCategoryUpdate{
			ID:          impactId,
			Name:        impactRequest.Name,
			ImpactPoint: impactRequest.ImpactPoint,
			IconURL:     impactRequest.IconURL,
		}

		err = h.i.Update(impact)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ImpactCategorySuccessUpdate, nil))
	}
}

func (h *ImpactCategoryHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
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

		impactId := c.Param("id")
		var impactCategory impactcategory.ImpactCategory
		impactCategory.ID = impactId
		err = h.i.Delete(impactCategory)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.ImpactCategorySuccessDelete, nil))
	}
}
