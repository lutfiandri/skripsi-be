package main

import (
	"log"

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

	userRepository := repository.NewUserRepository(mongo, "users")
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(app, authService)
	authController.InitHttpRoute()

	deviceTypeRepository := repository.NewDeviceTypeRepository(mongo, "device_types")
	deviceTypeService := service.NewDeviceTypeService(deviceTypeRepository)
	deviceTypeController := controller.NewDeviceTypeController(app, deviceTypeService)
	deviceTypeController.InitHttpRoute()

	deviceRepository := repository.NewDeviceRepository(mongo, "devices")
	deviceService := service.NewDeviceService(deviceRepository)
	deviceController := controller.NewDeviceController(app, deviceService)
	deviceController.InitHttpRoute()

	profileService := service.NewProfileService(userRepository)
	profileController := controller.NewProfileController(app, profileService)
	profileController.InitHttpRoute()

	app.Listen(":8080")
}
