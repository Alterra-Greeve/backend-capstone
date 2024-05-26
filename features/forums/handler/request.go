package handler

type CreateForumRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type EditForumRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
type CreateMessageForumRequest struct {
	ForumID  string `json:"forum_id" `
	Messages string `json:"messages" `
}

type EditMessageRequest struct {
	ForumID  string `json:"forum_id" `
	Messages string `json:"messages" `
}
