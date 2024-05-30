package handler

type UserbyAdminandPageResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Coin      int    `json:"coin"`
	Exp       int    `json:"exp"`
	AvatarURL string `json:"avatar_url"`
	Page      int    `json:"page"`
}

type UserbyAdminResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Coin      int    `json:"coin"`
	Exp       int    `json:"exp"`
	AvatarURL string `json:"avatar_url"`
}
