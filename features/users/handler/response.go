package handler

type UserInfoResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
	Phone       string `json:"phone"`
	Coin        int    `json:"coin"`
	Exp         int    `json:"exp"`
	ImpactPoint int    `json:"impact_point"`
	AvatarURL   string `json:"avatar_url"`
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

// Admin
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
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type MetadataResponse struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
}

type UserLeaderboardResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Exp       int    `json:"exp"`
	AvatarURL string `json:"avatar_url"`
	Rank      int    `json:"rank"`
}
