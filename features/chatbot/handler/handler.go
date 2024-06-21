package handler

import (
	"backendgreeve/constant"
	"backendgreeve/features/chatbot"
	"backendgreeve/helper"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ChatbotHandler struct {
	s chatbot.ChatbotServiceInterface
	j helper.JWTInterface
}

func New(s chatbot.ChatbotServiceInterface, j helper.JWTInterface) chatbot.ChatbotHandlerInterface {
	return &ChatbotHandler{
		s: s,
		j: j,
	}
}

func (ch *ChatbotHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := ch.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}
		var chatBotRequest ChatbotRequest
		if err := c.Bind(&chatBotRequest); err != nil {
			code, message := helper.HandleEchoError(err)
			return c.JSON(code, helper.FormatResponse(false, message, nil))
		}

		chatBot := chatbot.Chatbot{
			ChatID:  chatBotRequest.ID,
			Message: chatBotRequest.Message,
		}

		res, err := ch.s.Create(chatBot)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}

		var chatBotResponse ChatbotResponse
		chatBotResponse.ID = res.ID
		chatBotResponse.ChatID = res.ChatID
		chatBotResponse.Role = res.Role
		chatBotResponse.Message = res.Message
		chatBotResponse.CreatedAt = res.CreatedAt.Format(time.DateTime)

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "Success Create Chatbot", chatBotResponse))
	}
}

func (ch *ChatbotHandler) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		token, err := ch.j.ValidateToken(tokenString)
		if err != nil {
			return helper.UnauthorizedError(c)
		}

		userData := ch.j.ExtractUserToken(token)
		userID := userData[constant.JWT_ID].(string)

		res, err := ch.s.GetByID(userID)
		if err != nil {
			return c.JSON(helper.ConvertResponseCode(err), helper.FormatResponse(false, err.Error(), nil))
		}
		var chatBotResponse []ChatbotResponse
		for _, v := range res {
			chatBotResponse = append(chatBotResponse, ChatbotResponse{
				ID:        v.ID,
				ChatID:    v.ChatID,
				Role:      v.Role,
				Message:   v.Message,
				CreatedAt: v.CreatedAt.Format(time.DateTime),
			})
		}
		return c.JSON(http.StatusOK, helper.FormatResponse(true, "Success Get Chatbot", chatBotResponse))
	}
}
