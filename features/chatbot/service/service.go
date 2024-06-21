package service

import (
	"backendgreeve/config"
	"backendgreeve/features/chatbot"
	"context"
	"log"

	"github.com/google/uuid"
	openai "github.com/sashabaranov/go-openai"
)

type ChatbotService struct {
	d         chatbot.ChatbotDataInterface
	openaicfg config.OpenAi
}

func New(d chatbot.ChatbotDataInterface, openaicfg config.OpenAi) chatbot.ChatbotServiceInterface {
	return &ChatbotService{
		d:         d,
		openaicfg: openaicfg,
	}
}

func (s *ChatbotService) Create(chatbots chatbot.Chatbot) (chatbot.Chatbot, error) {
	var systemChatbot chatbot.Message
	systemChatbot.Role = "system"
	systemChatbot.Message = "Anda adalah seorang ahli dalam penghijauan lingkungan, anda akan diberikan pertanyaan tentang lingkungan dan anda harus menjawab pertanyaan tersebut dengan berhubungan tentang lingkungan, user kebanyakan akan mengirimkan pertanyaan tentang produk eco friendly dan challenge yang membuat lingkungan hijau."

	var payloadChatbot []chatbot.Message
	var ID string

	// Always add the system message first
	payloadChatbot = append(payloadChatbot, systemChatbot)

	if chatbots.ChatID == "" {
		// Generate a new chat ID for a new session
		ID = uuid.New().String()
		log.Println("sini 1")
	} else {
		// Use the provided chat ID and fetch the chat history
		log.Println("sini 2")
		ID = chatbots.ChatID
		res, err := s.d.GetByID(ID)
		if err != nil {
			return chatbot.Chatbot{}, err
		}
		for _, v := range res {
			payloadChatbot = append(payloadChatbot, chatbot.Message{
				Role:    v.Role,
				Message: v.Message,
			})
		}
	}
	// Create the user's message in the database
	chatbots.ID = uuid.New().String()
	chatbots.ChatID = ID
	chatbots.Role = "user"

	_, err := s.d.Create(chatbots)
	if err != nil {
		return chatbot.Chatbot{}, err
	}

	// Append the user's current message to the payload
	payloadChatbot = append(payloadChatbot, chatbot.Message{
		Role:    "user",
		Message: chatbots.Message,
	})

	// Convert the payload messages to the format required by the OpenAI API
	var openAIPayload []openai.ChatCompletionMessage
	for _, msg := range payloadChatbot {
		openAIPayload = append(openAIPayload, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Message,
		})
	}
	log.Println(openAIPayload)
	client := openai.NewClient(s.openaicfg.ApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4,
			Messages: openAIPayload,
		},
	)

	if err != nil {
		return chatbot.Chatbot{}, err
	}

	// Create the assistant's response in the database
	var chatBotAssistant chatbot.Chatbot
	chatBotAssistant.ID = uuid.New().String()
	chatBotAssistant.ChatID = ID
	chatBotAssistant.Role = "assistant"
	chatBotAssistant.Message = resp.Choices[0].Message.Content

	res, err := s.d.Create(chatBotAssistant)
	if err != nil {
		return chatbot.Chatbot{}, err
	}

	return res, nil
}

func (s *ChatbotService) GetByID(id string) ([]chatbot.Chatbot, error) {
	res, err := s.d.GetByID(id)
	if err != nil {
		return []chatbot.Chatbot{}, err
	}
	return res, nil
}
