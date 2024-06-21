package data

import (
	"backendgreeve/features/chatbot"

	"gorm.io/gorm"
)

type ChatbotData struct {
	DB *gorm.DB
}

func New(db *gorm.DB) chatbot.ChatbotDataInterface {
	return &ChatbotData{
		DB: db,
	}
}

func (cd *ChatbotData) Create(chatbots chatbot.Chatbot) (chatbot.Chatbot, error) {
	result := cd.DB.Create(&chatbots)
	if result.Error != nil {
		return chatbot.Chatbot{}, result.Error
	}
	return chatbots, nil
}

func (cd *ChatbotData) GetByID(id string) ([]chatbot.Chatbot, error) {
	var chatbots []chatbot.Chatbot
	result := cd.DB.Where("chat_id = ?", id).
		Order("created_at DESC").
		Find(&chatbots)
	if result.Error != nil {
		return []chatbot.Chatbot{}, result.Error
	}
	return chatbots, nil
}
