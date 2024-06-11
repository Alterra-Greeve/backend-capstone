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
			Stock:       productInput.Stock,
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

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.ProductSuccessCreate, nil))
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
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
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
					ImpactCategory: ImpactCategory{
						Name:        cat.ImpactCategory.Name,
						ImpactPoint: cat.ImpactCategory.ImpactPoint,
						IconURL:     cat.ImpactCategory.IconURL,
					},
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
				Stock:       p.Stock,
				CreatedAt:   p.CreatedAt.Format("02/01/2006"),
				UpdatedAt:   p.UpdatedAt.Format("02/01/2006"),
			})
		}
		metadata := MetadataResponse{
			TotalPage: totalPages,
			Page:      page,
		}

		return c.JSON(http.StatusOK, helper.MetadataFormatResponse(true, constant.ProductSuccessGet, metadata, response))
	}
}

func (h *ProductHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
			return nil
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
			return nil
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID]
		userRole := userData[constant.JWT_ROLE]

		paramId := c.Param("id")
		productId, err := uuid.Parse(paramId)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}
		var product product.Product
		if userRole == constant.RoleAdmin {
			product, err = h.s.GetById(productId.String())
			if err != nil {
				return c.JSON(helper.ConvertResponseCode(constant.ErrGetProductID), helper.FormatResponse(false, constant.ErrGetProductID.Error(), nil))
			}
		} else {
			product, err = h.s.GetByIdUser(productId.String(), userId.(string))
			if err != nil {
				return c.JSON(helper.ConvertResponseCode(constant.ErrGetUser), helper.FormatResponse(false, constant.ErrGetUser.Error(), nil))
			}
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
				ImpactCategory: ImpactCategory{
					Name:        cat.ImpactCategory.Name,
					ImpactPoint: cat.ImpactCategory.ImpactPoint,
					IconURL:     cat.ImpactCategory.IconURL,
				},
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
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.Format("02/01/2006"),
			UpdatedAt:   product.UpdatedAt.Format("02/01/2006"),
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ProductSuccessGet, response))
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
		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		var totalPages int
		var products []product.Product
		products, totalPages, err = h.s.GetByCategory(productCategory, page)
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
					ImpactCategory: ImpactCategory{
						Name:        cat.ImpactCategory.Name,
						ImpactPoint: cat.ImpactCategory.ImpactPoint,
						IconURL:     cat.ImpactCategory.IconURL,
					},
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
				Stock:       p.Stock,
				CreatedAt:   p.CreatedAt.Format("02/01/2006"),
				UpdatedAt:   p.UpdatedAt.Format("02/01/2006"),
			})
		}
		metadata := MetadataResponse{
			TotalPage: totalPages,
			Page:      page,
		}
		return c.JSON(http.StatusOK, helper.MetadataFormatResponse(true, constant.ProductSuccessGet, metadata, response))
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
		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		var totalPages int
		var products []product.Product
		products, totalPages, err = h.s.GetByName(productName, page)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
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
					ImpactCategory: ImpactCategory{
						Name:        cat.ImpactCategory.Name,
						ImpactPoint: cat.ImpactCategory.ImpactPoint,
						IconURL:     cat.ImpactCategory.IconURL,
					},
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
				Stock:       p.Stock,
				CreatedAt:   p.CreatedAt.Format("02/01/2006"),
				UpdatedAt:   p.UpdatedAt.Format("02/01/2006"),
			})
		}
		metadata := MetadataResponse{
			TotalPage: totalPages,
			Page:      page,
		}
		return c.JSON(http.StatusOK, helper.MetadataFormatResponse(true, constant.ProductSuccessGet, metadata, response))
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
			Stock:       productInput.Stock,
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

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, constant.ProductSuccessUpdate, nil))
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
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, constant.ProductSuccessDelete, nil))
	}
}
