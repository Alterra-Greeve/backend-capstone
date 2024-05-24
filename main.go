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

	ImpactCategoryData "backendgreeve/features/impactcategory/data"
	ImpactCategoryHandler "backendgreeve/features/impactcategory/handler"
	ImpactCategoryService "backendgreeve/features/impactcategory/service"

	"backendgreeve/helper"
	"backendgreeve/routes"
	"backendgreeve/utils/bucket"
	"backendgreeve/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.InitConfig()
	db, err := database.InitDB(*cfg)
	if err != nil {
		logrus.Error("terjadi kesalahan pada database, error:", err.Error())
	}
	database.Migrate(db)
	mailer := helper.NewMailer(cfg.SMTP)
	jwt := helper.NewJWT(cfg.JWT_Secret)
	bucket, _ := bucket.New(cfg.PROJECT_ID, cfg.BUCKET_NAME)
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})
	e.Static("/assets", "assets")
	e.Static("/docs", "docs")

	userData := UserData.New(db)
	userService := UserService.New(userData, jwt, mailer)
	userHandler := UserHandler.New(userService, jwt)

	webhookData := WebhookMidtransData.New(db)
	webhookService := WebhookMidtransService.New(webhookData)
	webhookHandler := WebhookMidtransHandler.New(webhookService)

	adminData := AdminData.New(db)
	adminService := AdminService.New(adminData, jwt, mailer)
	adminHandler := AdminHandler.New(adminService, jwt)

	impactCategoryData := ImpactCategoryData.New(db)
	impactCategoryService := ImpactCategoryService.New(impactCategoryData)
	impactCategoryHandler := ImpactCategoryHandler.New(impactCategoryService, jwt)

	routes.RouteUser(e, userHandler, *cfg)
	routes.RouteBucket(e, bucket, *cfg)
	routes.PaymentNotification(e, webhookHandler, *cfg)
	routes.RouteAdmin(e, adminHandler, *cfg)
	routes.RouteImpactCategory(e, impactCategoryHandler, *cfg)
	e.Logger.Fatal(e.Start(":8080"))
}
