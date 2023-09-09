package main

import (
	"skripsi-be/internal/controller"
	"skripsi-be/internal/infrastructure"
	"skripsi-be/internal/middleware"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	mongo := infrastructure.NewMongoDatabase("mongodb://root:root@localhost:27017", "skripsi-be")

	appConfig := fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	}
	app := fiber.New(appConfig)

	deviceTypeRepository := repository.NewDeviceTypeRepository(mongo, "device-types")
	deviceTypeService := service.NewDeviceTypeService(deviceTypeRepository)
	deviceTypeController := controller.NewDeviceTypeController(app, deviceTypeService)
	deviceTypeController.InitHttpRoute()

	app.Listen(":8080")
}
