package main

import (
	"backendgreeve/config"
	UserData "backendgreeve/features/users/data"
	UserHandler "backendgreeve/features/users/handler"
	UserService "backendgreeve/features/users/service"

	WebhookMidtransData "backendgreeve/features/webhook/data"
	WebhookMidtransHandler "backendgreeve/features/webhook/handler"
	WebhookMidtransService "backendgreeve/features/webhook/service"

	AdminData "backendgreeve/features/admin/data"
	AdminHandler "backendgreeve/features/admin/handler"
	AdminService "backendgreeve/features/admin/service"

	UserDatabyAdmin "backendgreeve/features/adminusers/data"
	usersAdminHandler "backendgreeve/features/adminusers/handler"
	UsersbyAdminService "backendgreeve/features/adminusers/service"

	ProductData "backendgreeve/features/product/data"
	ProductHandler "backendgreeve/features/product/handler"
	ProductService "backendgreeve/features/product/service"

	ImpactCategoryData "backendgreeve/features/impactcategory/data"
	ImpactCategoryHandler "backendgreeve/features/impactcategory/handler"
	ImpactCategoryService "backendgreeve/features/impactcategory/service"

	ForumData "backendgreeve/features/forums/data"
	ForumHandler "backendgreeve/features/forums/handler"
	ForumService "backendgreeve/features/forums/service"

	"backendgreeve/helper"
	"backendgreeve/routes"
	"backendgreeve/utils/bucket"
	"backendgreeve/utils/database"
	"backendgreeve/utils/database/seeds"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.InitConfig()
	db, err := database.InitDB(*cfg)
	if err != nil {
		logrus.Error("terjadi kesalahan pada database, error:", err.Error())
	}
	database.Migrate(db)

	for _, seed := range seeds.Seeds() {
		if err := seed.Run(db); err != nil {
			logrus.Error("terjadi kesalahan pada seed "+seed.Name+", error:", err.Error())
		}
	}

	mailer := helper.NewMailer(cfg.SMTP)
	jwt := helper.NewJWT(cfg.JWT_Secret)
	bucket, _ := bucket.New(cfg.PROJECT_ID, cfg.BUCKET_NAME)
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.Static("/assets", "assets")
	e.Static("/docs", "docs")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	userData := UserData.New(db)
	userService := UserService.New(userData, jwt, mailer)
	userHandler := UserHandler.New(userService, jwt)

	webhookData := WebhookMidtransData.New(db)
	webhookService := WebhookMidtransService.New(webhookData)
	webhookHandler := WebhookMidtransHandler.New(webhookService)

	adminData := AdminData.New(db)
	adminService := AdminService.New(adminData, jwt, mailer)
	adminHandler := AdminHandler.New(adminService, jwt)

	UserDatabyAdmin := UserDatabyAdmin.New(db)
	UsersbyAdminService := UsersbyAdminService.New(UserDatabyAdmin, jwt, mailer)
	UsersbyAdminHandler := usersAdminHandler.New(UsersbyAdminService, jwt)

	productData := ProductData.New(db)
	productService := ProductService.New(productData)
	productHandler := ProductHandler.New(productService, jwt)

	impactCategoryData := ImpactCategoryData.New(db)
	impactCategoryService := ImpactCategoryService.New(impactCategoryData)
	impactCategoryHandler := ImpactCategoryHandler.New(impactCategoryService, jwt)

	forumData := ForumData.New(db)
	forumService := ForumService.New(forumData)
	forumHandler := ForumHandler.New(forumService, jwt)

	routes.RouteUser(e, userHandler, *cfg)
	routes.RouteBucket(e, bucket, *cfg)
	routes.PaymentNotification(e, webhookHandler, *cfg)
	routes.RouteAdmin(e, adminHandler, *cfg)
	routes.RouteAdminManageUser(e, UsersbyAdminHandler, *cfg)
	routes.RouteProduct(e, productHandler, *cfg)
	routes.RouteImpactCategory(e, impactCategoryHandler, *cfg)
	routes.RouteForum(e, forumHandler, *cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
