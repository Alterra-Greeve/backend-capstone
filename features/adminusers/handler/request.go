package handler

type UserbyAdminRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
}
