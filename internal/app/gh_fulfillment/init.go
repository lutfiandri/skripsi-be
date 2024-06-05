package gh_fulfillment

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(app *fiber.App, database *mongo.Database, mqttClient mqtt.Client) {
	repository := NewRepository(database)
	service := NewService(repository, mqttClient)
	controller := NewController(service)
	Route(app, controller)
}
