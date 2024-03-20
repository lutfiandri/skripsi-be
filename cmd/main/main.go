package main

import (
	"log"

	"skripsi-be/internal/app/auth"
	"skripsi-be/internal/app/device"
	"skripsi-be/internal/app/device_type"
	"skripsi-be/internal/app/oauth_client"
	"skripsi-be/internal/app/oauth_scope"
	"skripsi-be/internal/app/profile"
	"skripsi-be/internal/config"
	"skripsi-be/internal/controller"
	"skripsi-be/internal/infrastructure"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("main service")

	mongo := infrastructure.NewMongoDatabase(config.MongoUri, config.MongoDbName)

	appConfig := fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	}
	app := fiber.New(appConfig)

	auth.Init(app, mongo)
	device_type.Init(app, mongo)
	device.Init(app, mongo)
	profile.Init(app, mongo)
	oauth_client.Init(app, mongo)
	oauth_scope.Init(app, mongo)

	oauthAuthCodeRepository := repository.NewOAuthAuthCodeRepository(mongo, "oauth_auth")

	oauthService := service.NewOAuthService(oauthClientRepository, oauthAuthCodeRepository, userRepository)
	oauthController := controller.NewOAuthController(app, oauthService)
	oauthController.InitHttpRoute()

	app.Listen(":8080")
}
