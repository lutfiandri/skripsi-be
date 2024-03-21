package oauth_scope

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(app *fiber.App, database *mongo.Database) {
	repository := NewRepository(database)
	service := NewService(repository)
	controller := NewController(app, service)
	Route(app, controller)
}