package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/product"
	"backendgreeve/helper"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	s product.ProductServiceInterface
	j helper.JWTInterface
}

func New(s product.ProductServiceInterface, j helper.JWTInterface) product.ProductHandlerInterface {
	return &ProductHandler{
		s: s,
		j: j,
	}
}

func (h *ProductHandler) Create() echo.HandlerFunc {
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

		var productInput ProductRequest
		err = c.Bind(&productInput)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		productData := product.Product{
			ID:          uuid.New().String(),
			Name:        productInput.Name,
			Description: productInput.Description,
			Price:       productInput.Price,
			Coin:        productInput.Coin,
		}

		for _, categoryID := range productInput.Category {
			productData.ImpactCategories = append(productData.ImpactCategories, product.ProductImpactCategory{
				ID:               uuid.New().String(),
				ProductID:        productData.ID,
				ImpactCategoryID: categoryID,
			})
		}

		for i, imageURL := range productInput.ImageURL {
			productData.Images = append(productData.Images, product.ProductImage{
				ID:        uuid.New().String(),
				ProductID: productData.ID,
				ImageURL:  imageURL,
				Position:  i + 1,
			})
		}

		err = h.s.Create(productData)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Product created successfully", nil))
	}
}

func (h *ProductHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
			return nil
		}

		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
			return nil
		}

		var products []product.Product
		products, err = h.s.Get()
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		var response []ProductResponse
		for _, p := range products {
			images := make([]ProductImage, len(p.Images))
			for i, img := range p.Images {
				images[i] = ProductImage{
					ID:        img.ID,
					ProductID: img.ProductID,
					ImageURL:  img.ImageURL,
					Position:  img.Position,
				}
			}

			categories := make([]ProductImpactCategory, len(p.ImpactCategories))
			for i, cat := range p.ImpactCategories {
				categories[i] = ProductImpactCategory{
					ID:               cat.ID,
					ProductID:        cat.ProductID,
					ImpactCategoryID: cat.ImpactCategoryID,
				}
			}

			response = append(response, ProductResponse{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Coin:        p.Coin,
				Images:      images,
				Category:    categories,
			})
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Product fetched successfully", []interface{}{response}))
	}
}

func (h *ProductHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
func (h *ProductHandler) GetByCategoryID() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
func (h *ProductHandler) GetByName() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
func (h *ProductHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
func (h *ProductHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
