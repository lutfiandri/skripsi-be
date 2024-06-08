package main

import (
	"context"
	"log"

	"skripsi-be/internal/app/auth"
	"skripsi-be/internal/app/device"
	"skripsi-be/internal/app/device_type"
	"skripsi-be/internal/app/gh_fulfillment"
	"skripsi-be/internal/app/oauth"
	"skripsi-be/internal/app/oauth_client"
	"skripsi-be/internal/app/oauth_scope"
	"skripsi-be/internal/app/permission"
	"skripsi-be/internal/app/profile"
	"skripsi-be/internal/app/role"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/util/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"google.golang.org/api/homegraph/v1"
)

func main() {
	log.Println("main service")

	mongo := infrastructure.NewMongoDatabase(config.MongoUri, config.MongoDbName)
	mqttClient := infrastructure.NewMqttClient(config.MqttBrokerUri, config.MqttUsername, config.MqttPassword)

	homegraphService, err := homegraph.NewService(context.Background())
	helper.PanicIfErr(err)
	homegraphDeviceService := homegraph.NewDevicesService(homegraphService)

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
	device.Init(app, mongo, mqttClient, homegraphDeviceService)
	profile.Init(app, mongo)
	oauth_client.Init(app, mongo)
	oauth_scope.Init(app, mongo)
	oauth.Init(app, mongo)
	role.Init(app, mongo)
	permission.Init(app, mongo)

	gh_fulfillment.Init(app, mongo, mqttClient)

	app.Listen(":6000")
}
