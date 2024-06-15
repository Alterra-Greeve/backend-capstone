package handler

import (
	"backendgreeve/constant"
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
				Difficulty:  challenge.Difficulty,
				Description: challenge.Description,
				Exp:         challenge.Exp,
				Coin:        challenge.Coin,
				ImageURL:    challenge.ImageURL,
				DateStart:   challenge.DateStart.Format("2006-01-02"),
				DateEnd:     challenge.DateEnd.Format("2006-01-02"),
			}

			var categories []ChallengeImpactCategories
			for _, cat := range challenge.ImpactCategories {
				categories = append(categories, ChallengeImpactCategories{
					ImpactCategory: ImpactCategory{
						Name:        cat.ImpactCategory.Name,
						ImpactPoint: cat.ImpactCategory.ImpactPoint,
						IconURL:     cat.ImpactCategory.IconURL,
					},
				})
			}
			challengeResponse.Categories = categories
			participant, err := h.s.GetChallengeParticipant(challenge.ID)
			if err != nil {
				return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
			}
			challengeResponse.Participant = participant
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
					ID:          cat.ImpactCategory.ID,
					Name:        cat.ImpactCategory.Name,
					ImpactPoint: cat.ImpactCategory.ImpactPoint,
					IconURL:     cat.ImpactCategory.IconURL,
				},
			}
		}
		response := ChallengeResponse{
			ID:          challenge.ID,
			Title:       challenge.Title,
			Description: challenge.Description,
			ImageURL:    challenge.ImageURL,
			Difficulty:  challenge.Difficulty,
			Exp:         challenge.Exp,
			Coin:        challenge.Coin,
			DateStart:   challenge.DateStart.Format("02/01/06"),
			DateEnd:     challenge.DateEnd.Format("02/01/06"),
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
				Difficulty:  challenge.Difficulty,
				Exp:         challenge.Exp,
				Coin:        challenge.Coin,
				DateStart:   challenge.DateStart.Format("02/01/06"),
				DateEnd:     challenge.DateEnd.Format("02/01/06"),
			}

			var categories []ChallengeImpactCategories
			for _, cat := range challenge.ImpactCategories {
				categories = append(categories, ChallengeImpactCategories{
					ImpactCategory: ImpactCategory{
						ID:          cat.ImpactCategory.ID,
						Name:        cat.ImpactCategory.Name,
						ImpactPoint: cat.ImpactCategory.ImpactPoint,
						IconURL:     cat.ImpactCategory.IconURL,
					},
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

func (h *ChallengeHandler) GetUserParticipate() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)
		status := c.QueryParam("status")

		challenges, err := h.s.GetUserParticipate(userId, status)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var response []UserChallengeConfirmationResponse
		for _, challenge := range challenges {
			response = append(response, new(UserChallengeConfirmationResponse).ToResponse(challenge))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success", response))
	}
}

func (h *ChallengeHandler) Update() echo.HandlerFunc {
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
		adminRole := adminData[constant.JWT_ROLE].(string)
		if adminRole != constant.RoleAdmin {
			helper.UnauthorizedError(c)
		}

		challengeId := c.Param("id")
		var request ChallengeCreateRequest
		if err := c.Bind(&request); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		challenge := challenges.Challenge{
			ID:          challengeId,
			Title:       request.Title,
			Description: request.Description,
			Difficulty:  request.Difficulty,
			Exp:         request.Exp,
			Coin:        request.Coin,
			ImageURL:    request.ImageURL,
			DateStart:   request.DateStart,
			DateEnd:     request.DateEnd,
		}

		for _, category := range request.ImpactCategories {
			challenge.ImpactCategories = append(challenge.ImpactCategories, challenges.ChallengeImpactCategory{
				ImpactCategoryID: category,
			})
		}
		err = h.s.Update(challenge)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success", nil))
	}
}

func (h *ChallengeHandler) Delete() echo.HandlerFunc {
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
		adminRole := adminData[constant.JWT_ROLE].(string)
		if adminRole != constant.RoleAdmin {
			helper.UnauthorizedError(c)
		}

		challengeId := c.Param("id")
		err = h.s.Delete(challengeId)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success", nil))
	}
}

func (h *ChallengeHandler) GetChallengeForUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		_, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		challengeConfirmationId := c.Param("id")
		challenge, err := h.s.GetChallengeForUserByID(challengeConfirmationId)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.ObjectFormatResponse(false, constant.ErrGetForumByID.Error(), nil))
		}
		userChallengeResponse := new(UserChallengeConfirmationResponse)
		response := userChallengeResponse.ToResponse(challenge)

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ForumSuccessGetByID, response))
	}
}
func (h *ChallengeHandler) EditChallengeForUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get(constant.HeaderAuthorization)
		if tokenString == "" {
			return helper.UnauthorizedError(c)
		}

		token, err := h.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		userData := h.j.ExtractUserToken(token)
		userId := userData[constant.JWT_ID].(string)

		challengeID := c.Param("id")
		existingChallenge, err := h.s.GetChallengeForUserByID(challengeID)
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, constant.ErrForumNotFound.Error(), nil))
		}

		if existingChallenge.UserID != userId {
			return c.JSON(http.StatusForbidden, helper.FormatResponse(false, constant.UnatuhorizeForumAndMessage.Error(), nil))
		}

		var challenge EditChallengeConfirmationForUser
		if err := c.Bind(&challenge); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		err = h.s.EditChallengeForUserByID(challengeID, challenge.Image)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.ObjectFormatResponse(true, constant.ForumSuccessUpdate, nil))
	}
}
