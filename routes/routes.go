package routes

import (
	"backendgreeve/config"
	"backendgreeve/features/users"
	MidtransWebhook "backendgreeve/features/webhook"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg config.GreeveConfig) {

	e.POST("/register", uh.Register())
	e.POST("/login", uh.Login())
	e.POST("/forgot-password", uh.ForgotPassword())
	e.POST("/verify-otp", uh.VerifyOTP())
	e.POST("/reset-password", uh.ResetPassword())

	e.PUT("/update-user", uh.Update(), echojwt.JWT([]byte(cfg.JWT_Secret)))
	e.DELETE("/delete-user", uh.Delete(), echojwt.JWT([]byte(cfg.JWT_Secret)))
}

func PaymentNotification(e *echo.Echo, wh MidtransWebhook.MidtransNotificationHandler, cfg config.GreeveConfig) {
	e.POST("/midtrans-notification", wh.HandleNotification())
}
