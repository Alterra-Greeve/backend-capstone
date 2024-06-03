package handler

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserForgotPasswordRequest struct {
	Email string `json:"email"`
}

type UserVerifyOTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}

type UserUpdateRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Address   string `json:"address"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
}

type UserUpdateAvatarRequest struct {
	Image string `json:"image"`
}

type UserCheckOTPRequest struct {
	Email string `json:"email"`
	OTP   string `json:"otp"`
}
type UserResetPasswordRequest struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	ConfirmationPassword string `json:"confirmation_password"`
	OTP                  string `json:"otp"`
}

// Admin
type UserbyAdminRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
}
