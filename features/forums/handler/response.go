package handler

type ForumGetAllResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      Author `json:"author"`
}

type ForumGetDetailResponse struct {
	ID            string            `json:"id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	Author        Author            `json:"author"`
	ForumMessages []MessageResponse `json:"forum_messages"`
}

type Author struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type MessageResponse struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}
