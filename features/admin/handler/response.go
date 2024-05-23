package handler

type AdminLoginResponse struct {
	Token string `json:"token"`
}

type AdminUpdateResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
}

type AdminInfoResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
