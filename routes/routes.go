package routes

import (
	"backendgreeve/config"
	"backendgreeve/constant/route"
	"backendgreeve/features/admin"
	"backendgreeve/features/impactcategory"
	"backendgreeve/features/product"
	"backendgreeve/features/users"
	MidtransWebhook "backendgreeve/features/webhook"
	"backendgreeve/helper"
	"backendgreeve/utils/bucket"

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

func RouteAdmin(e *echo.Echo, ah admin.AdminHandlerInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}

	e.POST(route.AdminLogin, ah.Login())

	e.GET(route.AdminPath, ah.GetAdminData(), echojwt.WithConfig(jwtConfig))
	e.PUT(route.AdminPath, ah.Update(), echojwt.WithConfig(jwtConfig))
	e.DELETE(route.AdminPath, ah.Delete(), echojwt.WithConfig(jwtConfig))
}

func RouteProduct(e *echo.Echo, ph product.ProductHandlerInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}

	e.POST("/api/v1/products", ph.Create())
	e.GET("/api/v1/products", ph.Get(), echojwt.WithConfig(jwtConfig))
	e.GET("/api/v1/product/:id", ph.GetById(), echojwt.WithConfig(jwtConfig))
	e.PUT("/api/v1/product/:id", ph.Update(), echojwt.WithConfig(jwtConfig))
	e.DELETE("/api/v1/product/:id", ph.Delete(), echojwt.WithConfig(jwtConfig))
}
func RouteImpactCategory(e *echo.Echo, ic impactcategory.ImpactCategoryHandlerInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}

	e.POST(route.ImpactCategoryPath, ic.Create(), echojwt.WithConfig(jwtConfig))
	e.GET(route.ImpactCategoryPath, ic.GetAll(), echojwt.WithConfig(jwtConfig))
	e.GET(route.ImpactCategoryByID, ic.GetByID(), echojwt.WithConfig(jwtConfig))
	e.PUT(route.ImpactCategoryByID, ic.Update(), echojwt.WithConfig(jwtConfig))
	e.DELETE(route.ImpactCategoryByID, ic.Delete(), echojwt.WithConfig(jwtConfig))
}

func RouteBucket(e *echo.Echo, bh bucket.BucketInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}

	e.POST("/api/v1/media/upload", bh.UploadFileHandler(), echojwt.WithConfig(jwtConfig))
}
func PaymentNotification(e *echo.Echo, wh MidtransWebhook.MidtransNotificationHandler, cfg config.GreeveConfig) {
	e.POST("/midtrans-notification", wh.HandleNotification())
}
