package chatbot

import "github.com/labstack/echo/v4"

type Chatbot struct {
	ID      string
	Name    string
	Role    string
	Message string
}

type ChatbotHandlerInterface interface {
	Create() echo.HandlerFunc
}

type ChatbotServiceInterface interface {
	Create(chat Chatbot) (Chatbot, error)
	GetByID(id string) (Chatbot, error)
}

type ChatbotDataInterface interface {
	Create(chat Chatbot) (Chatbot, error)
	GetByID(id string) (Chatbot, error)
}