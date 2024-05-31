package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	JWTSecretKey        = os.Getenv("JWT_SECRET_KEY")
	JWTRefreshSecretKey = os.Getenv("JWT_REFRESH_SECRET_KEY")
	MongoUri            = os.Getenv("MONGO_URI")
	MongoDbName         = os.Getenv("MONGO_DB_NAME")
	MqttBrokerUri       = os.Getenv("MQTT_BROKER_URI")
	MqttUsername        = os.Getenv("MQTT_USERNAME")
	MqttPassword        = os.Getenv("MQTT_PASSWORD")
	KafkaBrokerUris     = os.Getenv("KAFKA_BROKER_URIS")
	RedisUri            = os.Getenv("REDIS_URI")
	SmtpEmail           = os.Getenv("SMTP_EMAIL")
	SmtpPassword        = os.Getenv("SMTP_PASSWORD")
	SmtpName            = os.Getenv("SMTP_NAME")
	SmtpHost            = os.Getenv("SMTP_HOST")
	SmtpPort            = os.Getenv("SMTP_PORT")

	KafkaTopicDeviceState = "device_state"
	KafkaGroupGhReporter  = "gh_reporter"
)
