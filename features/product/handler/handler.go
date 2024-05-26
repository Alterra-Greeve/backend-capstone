package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/product"
	"backendgreeve/helper"
	"net/http"
	"strconv"

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

		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		var totalPages int
		var products []product.Product
		products, totalPages, err = h.s.GetByPage(page)

		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		var response []ProductResponse
		for _, p := range products {
			images := make([]ProductImage, len(p.Images))
			for i, img := range p.Images {
				images[i] = ProductImage{
					ImageURL: img.ImageURL,
					Position: img.Position,
				}
			}

			categories := make([]ProductImpactCategory, len(p.ImpactCategories))
			for i, cat := range p.ImpactCategories {
				categories[i] = ProductImpactCategory{
					ImpactCategory: ImpactCategory{Name: cat.ImpactCategory.Name, ImpactPoint: cat.ImpactCategory.ImpactPoint},
				}
			}

			response = append(response, ProductResponse{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Coin:        p.Coin,
				CurrentPage: page,
				TotalPage:   totalPages,
				Images:      images,
				Category:    categories,
			})
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Product fetched successfully", []interface{}{response}))
	}
}

func (h *ProductHandler) GetById() echo.HandlerFunc {
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

		productID := c.Param("id")
		product, err := h.s.GetById(productID)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		images := make([]ProductImage, len(product.Images))
		for i, img := range product.Images {
			images[i] = ProductImage{
				ImageURL: img.ImageURL,
				Position: img.Position,
			}
		}

		categories := make([]ProductImpactCategory, len(product.ImpactCategories))
		for i, cat := range product.ImpactCategories {
			categories[i] = ProductImpactCategory{
				ImpactCategory: ImpactCategory{Name: cat.ImpactCategory.Name, ImpactPoint: cat.ImpactCategory.ImpactPoint},
			}
		}

		response := ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Coin:        product.Coin,
			Images:      images,
			Category:    categories,
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, "Product fetched successfully", response))
	}
}
func (h *ProductHandler) GetByCategory() echo.HandlerFunc {
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

		productCategory := c.Param("category_name")
		var products []product.Product
		products, err = h.s.GetByCategory(productCategory)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		var response []ProductResponse
		for _, p := range products {
			images := make([]ProductImage, len(p.Images))
			for i, img := range p.Images {
				images[i] = ProductImage{
					ImageURL: img.ImageURL,
					Position: img.Position,
				}
			}

			categories := make([]ProductImpactCategory, len(p.ImpactCategories))
			for i, cat := range p.ImpactCategories {
				categories[i] = ProductImpactCategory{
					ImpactCategory: ImpactCategory{Name: cat.ImpactCategory.Name, ImpactPoint: cat.ImpactCategory.ImpactPoint},
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
func (h *ProductHandler) GetByName() echo.HandlerFunc {
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

		productName := c.QueryParam("name")
		var products []product.Product
		products, err = h.s.GetByName(productName)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		var response []ProductResponse
		for _, p := range products {
			images := make([]ProductImage, len(p.Images))
			for i, img := range p.Images {
				images[i] = ProductImage{
					ImageURL: img.ImageURL,
					Position: img.Position,
				}
			}

			categories := make([]ProductImpactCategory, len(p.ImpactCategories))
			for i, cat := range p.ImpactCategories {
				categories[i] = ProductImpactCategory{
					ImpactCategory: ImpactCategory{Name: cat.ImpactCategory.Name, ImpactPoint: cat.ImpactCategory.ImpactPoint},
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
func (h *ProductHandler) Update() echo.HandlerFunc {
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

		productID := c.Param("id")
		var productInput ProductRequest
		err = c.Bind(&productInput)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		productData := product.Product{
			ID:          productID,
			Name:        productInput.Name,
			Description: productInput.Description,
			Price:       productInput.Price,
			Coin:        productInput.Coin,
		}

		for _, categoryID := range productInput.Category {
			productData.ImpactCategories = append(productData.ImpactCategories, product.ProductImpactCategory{
				ID:               uuid.New().String(),
				ProductID:        productID,
				ImpactCategoryID: categoryID,
			})
		}

		for i, imageURL := range productInput.ImageURL {
			productData.Images = append(productData.Images, product.ProductImage{
				ID:        uuid.New().String(),
				ProductID: productID,
				ImageURL:  imageURL,
				Position:  i + 1,
			})
		}

		err = h.s.Update(productData)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Product update successfully", nil))
	}
}

func (h *ProductHandler) Delete() echo.HandlerFunc {
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

		productID := c.Param("id")
		var productInput ProductRequest
		err = c.Bind(&productInput)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		err = h.s.Delete(productID)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Product deleted successfully", nil))
	}
}
