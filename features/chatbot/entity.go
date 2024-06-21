package chatbot

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Chatbot struct {
	ID        string
	ChatID    string
	Role      string
	Message   string
	CreatedAt time.Time
}

type Message struct {
	Role    string
	Message string
}
type ChatbotHandlerInterface interface {
	Create() echo.HandlerFunc
	GetByID() echo.HandlerFunc
}

type ChatbotServiceInterface interface {
	Create(chat Chatbot) (Chatbot, error)
	GetByID(id string) ([]Chatbot, error)
}

type ChatbotDataInterface interface {
	Create(chat Chatbot) (Chatbot, error)
	GetByID(id string) ([]Chatbot, error)
}
