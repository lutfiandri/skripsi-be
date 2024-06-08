package device

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/homegraph/v1"
)

func Init(app *fiber.App, database *mongo.Database, mqttClient mqtt.Client, homegraphDeviceService *homegraph.DevicesService) {
	repository := NewRepository(database)
	service := NewService(repository, mqttClient, homegraphDeviceService)
	controller := NewController(app, service)
	Route(app, controller)
}
