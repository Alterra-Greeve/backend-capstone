package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/challenges"
	"backendgreeve/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ChallengeHandler struct {
	s challenges.ChallengeServiceInterface
	j helper.JWTInterface
}

func New(s challenges.ChallengeServiceInterface, j helper.JWTInterface) challenges.ChallengeHandlerInterface {
	return &ChallengeHandler{
		s: s,
		j: j,
	}
}

func (h *ChallengeHandler) GetAllForUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)

		challenges, err := h.s.GetAllForUser(userId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		var response []ChallengeResponse
		for _, challenge := range challenges {
			challengeResponse := ChallengeResponse{
				ID:          challenge.ID,
				Title:       challenge.Title,
				Description: challenge.Description,
				ImageURL:    challenge.ImageURL,
				DateStart:   challenge.DateStart.Format("2006-01-02"),
				DateEnd:     challenge.DateEnd.Format("2006-01-02"),
			}

			var categories []ChallengeImpactCategories
			for _, cat := range challenge.ImpactCategories {
				categories = append(categories, ChallengeImpactCategories{
					ImpactCategory: ImpactCategory{Name: cat.ImpactCategory.Name, ImpactPoint: cat.ImpactCategory.ImpactPoint},
				})
			}
			challengeResponse.Categories = categories
			response = append(response, challengeResponse)
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, "Success", response))
	}
}

func (h *ChallengeHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		id := c.Param("id")
		challenge, err := h.s.GetByID(id)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		category := make([]ChallengeImpactCategories, len(challenge.ImpactCategories))
		for i, cat := range challenge.ImpactCategories {
			category[i] = ChallengeImpactCategories{
				ImpactCategory: ImpactCategory{
					Name:        cat.ImpactCategory.Name,
					ImpactPoint: cat.ImpactCategory.ImpactPoint,
				},
			}
		}
		response := ChallengeResponse{
			ID:          challenge.ID,
			Title:       challenge.Title,
			Description: challenge.Description,
			ImageURL:    challenge.ImageURL,
			DateStart:   challenge.DateStart.Format("2006-01-02"),
			DateEnd:     challenge.DateEnd.Format("2006-01-02"),
			Categories:  category,
		}
		response.Categories = category
		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, "Success", response))
	}
}

func (h *ChallengeHandler) Swipe() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)

		var challengeRequest ChallengeParticipateRequest
		if err := c.Bind(&challengeRequest); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		err = h.s.Swipe(userId, challengeRequest.ChallengeId, challengeRequest.ChallengeType)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success", nil))
	}
}

func (h *ChallengeHandler) GetAllForAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {

		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}

		challenges, totalPages, err := h.s.GetAllForAdmin(page)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response []ChallengeResponse
		for _, challenge := range challenges {
			challengeResponse := ChallengeResponse{
				ID:          challenge.ID,
				Title:       challenge.Title,
				Description: challenge.Description,
				ImageURL:    challenge.ImageURL,
				DateStart:   challenge.DateStart.Format("2006-01-02"),
				DateEnd:     challenge.DateEnd.Format("2006-01-02"),
			}

			var categories []ChallengeImpactCategories
			for _, cat := range challenge.ImpactCategories {
				categories = append(categories, ChallengeImpactCategories{
					ImpactCategory: ImpactCategory{Name: cat.ImpactCategory.Name, ImpactPoint: cat.ImpactCategory.ImpactPoint},
				})
			}

			challengeResponse.Categories = categories
			response = append(response, challengeResponse)
		}

		metadata := MetadataResponse{
			CurrentPage: page,
			TotalPage:   totalPages,
		}

		return c.JSON(http.StatusOK, helper.MetadataFormatResponse(true, "Success", metadata, response))
	}
}

func (h *ChallengeHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			helper.UnauthorizedError(c)
		}
		adminData := h.j.ExtractUserToken(token)
		adminRole := adminData[constant.JWT_ROLE].(string)
		if adminRole != constant.RoleAdmin {
			helper.UnauthorizedError(c)
		}

		var challengeRequest ChallengeCreateRequest
		if err := c.Bind(&challengeRequest); err != nil {
			log.Println(err)
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}
		challenge := challenges.Challenge{
			Title:       challengeRequest.Title,
			Difficulty:  challengeRequest.Difficulty,
			Description: challengeRequest.Description,
			Exp:         challengeRequest.Exp,
			Coin:        challengeRequest.Coin,
			ImageURL:    challengeRequest.ImageURL,
			DateStart:   challengeRequest.DateStart,
			DateEnd:     challengeRequest.DateEnd,
		}
		for _, category := range challengeRequest.ImpactCategories {
			challenge.ImpactCategories = append(challenge.ImpactCategories, challenges.ChallengeImpactCategory{
				ImpactCategoryID: category,
			})
		}
		err = h.s.Create(challenge)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Success", nil))
	}
}

func (h *ChallengeHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		return nil
	}
}

func (h *ChallengeHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		return nil
	}
}
