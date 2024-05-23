package routes

import (
	"backendgreeve/config"
	"backendgreeve/constant/route"
	"backendgreeve/features/users"
	MidtransWebhook "backendgreeve/features/webhook"
	"backendgreeve/helper"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uh users.UserHandlerInterface, cfg config.GreeveConfig) {

	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}

	e.POST(route.UserRegister, uh.Register())
	e.POST(route.UserLogin, uh.Login())
	e.POST(route.UserForgotPassword, uh.ForgotPassword())
	e.POST(route.UserVerifyOTP, uh.VerifyOTP())
	e.POST(route.UserResetPassword, uh.ResetPassword())

	e.GET(route.UserByUsername, uh.GetUserByUsername())
	e.GET(route.UserPath, uh.GetUserData(), echojwt.WithConfig(jwtConfig))
	e.PUT(route.UserPath, uh.Update(), echojwt.WithConfig(jwtConfig))
	e.DELETE(route.UserPath, uh.Delete(), echojwt.WithConfig(jwtConfig))
}

func PaymentNotification(e *echo.Echo, wh MidtransWebhook.MidtransNotificationHandler, cfg config.GreeveConfig) {
	e.POST("/midtrans-notification", wh.HandleNotification())
}
