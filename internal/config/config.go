package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	JWTSecretKey  = os.Getenv("JWT_SECRET_KEY")
	MongoUri      = os.Getenv("MONGO_URI")
	MongoDbName   = os.Getenv("MONGO_DB_NAME")
	MqttBrokerUri = os.Getenv("MQTT_BROKER_URI")
	MqttUsername  = os.Getenv("MQTT_USERNAME")
	MqttPassword  = os.Getenv("MQTT_PASSWORD")
)
