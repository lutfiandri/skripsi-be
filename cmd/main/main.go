package main

import (
	"log"

	"skripsi-be/internal/app/auth"
	"skripsi-be/internal/app/device"
	"skripsi-be/internal/app/device_type"
	"skripsi-be/internal/app/gh_fulfillment"
	"skripsi-be/internal/app/oauth"
	"skripsi-be/internal/app/oauth_client"
	"skripsi-be/internal/app/oauth_scope"
	"skripsi-be/internal/app/profile"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
	"skripsi-be/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	log.Println("main service")

	mongo := infrastructure.NewMongoDatabase(config.MongoUri, config.MongoDbName)

	appConfig := fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	}
	app := fiber.New(appConfig)
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(logger.New())

	auth.Init(app, mongo)
	device_type.Init(app, mongo)
	device.Init(app, mongo)
	profile.Init(app, mongo)
	oauth_client.Init(app, mongo)
	oauth_scope.Init(app, mongo)
	oauth.Init(app, mongo)

	gh_fulfillment.Init(app, mongo)

	app.Listen(":6000")
}
