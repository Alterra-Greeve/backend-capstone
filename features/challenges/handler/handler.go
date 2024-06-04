package handler

import (
	"backendgreeve/features/challenges"
	"backendgreeve/helper"
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
		return nil
	}
}

func (h *ChallengeHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		return nil
	}
}

func (h *ChallengeHandler) Swipe() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Your code here
		return nil
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

			challengeResponse.ImpactCategories = categories
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
		return nil
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
