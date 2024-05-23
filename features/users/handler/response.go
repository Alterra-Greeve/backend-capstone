package handler

type UserInfoResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Coin      int    `json:"coin"`
	Exp       int    `json:"exp"`
	AvatarURL string `json:"avatar_url"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
type UserUpdateResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
}
