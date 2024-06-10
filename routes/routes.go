package routes

import (
	"backendgreeve/config"
	"backendgreeve/constant/route"
	"backendgreeve/features/admin"
	"backendgreeve/features/cart"
	"backendgreeve/features/challenges"
	"backendgreeve/features/forums"
	"backendgreeve/features/impactcategory"
	"backendgreeve/features/product"
	"backendgreeve/features/transaction"
	"backendgreeve/features/users"
	"backendgreeve/features/voucher"
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

	// Leaderboard
	e.GET(route.Leaderboard, uh.GetLeaderboard(), echojwt.WithConfig(jwtConfig))

	// Admin
	e.GET(route.AdminManageUserPath, uh.GetAllUsersForAdmin(), echojwt.WithConfig(jwtConfig))
	e.GET(route.AdminManageUserByID, uh.GetUserByIDForAdmin(), echojwt.WithConfig(jwtConfig))
	e.PUT(route.AdminManageUserByID, uh.UpdateUserForAdmin(), echojwt.WithConfig(jwtConfig))
	e.DELETE(route.AdminManageUserByID, uh.DeleteUserForAdmin(), echojwt.WithConfig(jwtConfig))
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
	e.GET("/api/v1/products/search", ph.GetByName(), echojwt.WithConfig(jwtConfig))
	e.GET("/api/v1/products/category/:category_name", ph.GetByCategory(), echojwt.WithConfig(jwtConfig))
	e.GET("/api/v1/products/:id", ph.GetById(), echojwt.WithConfig(jwtConfig))
	e.PUT("/api/v1/admin/products/:id", ph.Update(), echojwt.WithConfig(jwtConfig))
	e.DELETE("/api/v1/admin/products/:id", ph.Delete(), echojwt.WithConfig(jwtConfig))
}

func RouteCart(e *echo.Echo, ch cart.CartHandlerInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}
	e.POST("/api/v1/cart", ch.Create(), echojwt.WithConfig(jwtConfig))
	e.GET("/api/v1/cart", ch.Get(), echojwt.WithConfig(jwtConfig))
	e.PUT("/api/v1/cart", ch.Update(), echojwt.WithConfig(jwtConfig))
	e.DELETE("/api/v1/cart", ch.Delete(), echojwt.WithConfig(jwtConfig))
}

func RouteTransaction(e *echo.Echo, th transaction.TransactionHandlerInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}
	e.GET("/api/v1/transactions", th.GetUserTransaction(), echojwt.WithConfig(jwtConfig))
	e.POST("/api/v1/transactions", th.CreateTransaction(), echojwt.WithConfig(jwtConfig))
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

func RouteVoucher(e *echo.Echo, vc voucher.VoucherHandlerInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}
	e.GET(route.VoucherPath, vc.GetAll(), echojwt.WithConfig(jwtConfig))
	e.GET(route.VoucherByID, vc.GetByIdVoucher(), echojwt.WithConfig(jwtConfig))
	e.POST(route.VoucherPath, vc.Create(), echojwt.WithConfig(jwtConfig))
	e.PUT(route.VoucherByID, vc.Update(), echojwt.WithConfig(jwtConfig))
	e.DELETE(route.VoucherByID, vc.Delete(), echojwt.WithConfig(jwtConfig))
}

func RouteChallenge(e *echo.Echo, ch challenges.ChallengeHandlerInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}

	e.POST(route.AdminChallengePath, ch.Create(), echojwt.WithConfig(jwtConfig))
	e.POST(route.ChallengeParticipate, ch.Swipe(), echojwt.WithConfig(jwtConfig))
	e.GET(route.ChallengePath, ch.GetAllForUser(), echojwt.WithConfig(jwtConfig))
	e.GET(route.AdminChallengePath, ch.GetAllForAdmin(), echojwt.WithConfig(jwtConfig))
	e.GET(route.ChallengeByID, ch.GetByID(), echojwt.WithConfig(jwtConfig))
	e.GET(route.ChallengeParticipate, ch.GetUserParticipate(), echojwt.WithConfig(jwtConfig))
	e.GET(route.AdminChallengeByID, ch.GetByID(), echojwt.WithConfig(jwtConfig))
	e.PUT(route.ChallengeByID, ch.Update(), echojwt.WithConfig(jwtConfig))
	e.DELETE(route.ChallengeByID, ch.Delete(), echojwt.WithConfig(jwtConfig))
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

func RouteForum(e *echo.Echo, fh forums.ForumHandlerInterface, cfg config.GreeveConfig) {
	jwtConfig := echojwt.Config{
		SigningKey:   []byte(cfg.JWT_Secret),
		ErrorHandler: helper.JWTErrorHandler,
	}

	e.GET(route.ForumPath, fh.GetAllForum(), echojwt.WithConfig(jwtConfig))
	e.GET(route.ForumByID, fh.GetForumByID(), echojwt.WithConfig(jwtConfig))
	e.GET(route.GetForumByUserID, fh.GetForumByUserID(), echojwt.WithConfig(jwtConfig))
	e.POST(route.ForumPath, fh.PostForum(), echojwt.WithConfig(jwtConfig))
	e.PUT(route.ForumByID, fh.UpdateForum(), echojwt.WithConfig(jwtConfig))
	e.DELETE(route.ForumByID, fh.DeleteForum(), echojwt.WithConfig(jwtConfig))

	e.GET(route.ForumMessageByID, fh.GetMessageForumByID(), echojwt.WithConfig(jwtConfig))
	e.POST(route.ForumMessage, fh.PostMessageForum(), echojwt.WithConfig(jwtConfig))
	e.DELETE(route.ForumMessageByID, fh.DeleteMessageForum(), echojwt.WithConfig(jwtConfig))
	e.PUT(route.ForumMessageByID, fh.UpdateMessageForum(), echojwt.WithConfig(jwtConfig))
}
