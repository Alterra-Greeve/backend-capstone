package handler

type ChatbotResponse struct {
	ID        string `json:"id"`
	ChatID    string `json:"chat_id"`
	Role      string `json:"role"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}
