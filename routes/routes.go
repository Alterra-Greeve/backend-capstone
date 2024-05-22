package routes

import (
	"backendgreeve/config"
	"backendgreeve/constant/route"
	"backendgreeve/features/users"
	MidtransWebhook "backendgreeve/features/webhook"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg config.GreeveConfig) {

	e.POST(route.UserRegister, uh.Register())
	e.POST(route.UserLogin, uh.Login())
	e.POST(route.UserForgotPassword, uh.ForgotPassword())
	e.POST(route.UserVerifyOTP, uh.VerifyOTP())
	e.POST(route.UserResetPassword, uh.ResetPassword())

	e.PUT(route.UserPath, uh.Update(), echojwt.JWT([]byte(cfg.JWT_Secret)))
	e.DELETE(route.UserPath, uh.Delete(), echojwt.JWT([]byte(cfg.JWT_Secret)))
}

func PaymentNotification(e *echo.Echo, wh MidtransWebhook.MidtransNotificationHandler, cfg config.GreeveConfig) {
	e.POST("/midtrans-notification", wh.HandleNotification())
}
