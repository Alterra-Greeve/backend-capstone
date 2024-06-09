package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/cart"
	product "backendgreeve/features/product/handler"
	"backendgreeve/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	s cart.CartServiceInterface
	j helper.JWTInterface
}

func New(s cart.CartServiceInterface, j helper.JWTInterface) cart.CartHandlerInterface {
	return &CartHandler{
		s: s,
		j: j,
	}
}

func (h *CartHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		var cartRequest CreateCartRequest
		err := c.Bind(&cartRequest)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}
		tokenString := c.Request().Header.Get("Authorization")

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID]

		newCart := cart.NewCart{
			ProductID: cartRequest.ProductID,
			UserID:    userId.(string),
		}
		err = h.s.Create(newCart)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Cart created", nil))
	}
}

func (h *CartHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var cartRequest UpdateCartRequest
		err := c.Bind(&cartRequest)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}
		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID]

		newCart := cart.UpdateCart{
			ProductID: cartRequest.ProductID,
			UserID:    userId.(string),
			Type:      cartRequest.Type,
			Quantity:  cartRequest.Quantity,
		}
		if cartRequest.Quantity != 0 {
			newCart.Type = "qty"
			newCart.Quantity = cartRequest.Quantity
		}
		if cartRequest.Type != "" {
			newCart.Type = cartRequest.Type
		}

		err = h.s.Update(newCart)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Cart updated", nil))
	}
}

func (h *CartHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "created")
	}
}

func (h *CartHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}
		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID]
		carts, err := h.s.Get(userId.(string))
		if err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		var response CartResponse
		response.User = User{
			ID:       carts.User.ID,
			Username: carts.User.Username,
			Email:    carts.User.Email,
			Address:  carts.User.Address,
		}

		for _, cart := range carts.Items {
			var images []product.ProductImage
			var impactCategories []product.ProductImpactCategory

			for _, img := range cart.Product.Images {
				images = append(images, product.ProductImage{
					ImageURL: img.ImageURL,
					Position: img.Position,
				})
			}

			for _, impact := range cart.Product.ImpactCategories {
				impactCategories = append(impactCategories, product.ProductImpactCategory{
					ImpactCategory: product.ImpactCategory{
						Name:        impact.ImpactCategory.Name,
						ImpactPoint: impact.ImpactCategory.ImpactPoint,
						IconURL:     impact.ImpactCategory.IconURL,
					},
				})
			}

			response.Items = append(response.Items, CartItems{
				Quantity: cart.Quantity,
				Product: product.ProductResponse{
					ID:          cart.Product.ID,
					Name:        cart.Product.Name,
					Description: cart.Product.Description,
					Price:       cart.Product.Price,
					Coin:        cart.Product.Coin,
					Stock:       cart.Product.Stock,
					CreatedAt:   cart.Product.CreatedAt.Format("2006-01-02"),
					UpdatedAt:   cart.Product.UpdatedAt.Format("2006-01-02"),
					Images:      images,
					Category:    impactCategories,
				},
			})
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get all carts", response))
	}
}
