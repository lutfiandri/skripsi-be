package main

import (
	"skripsi-be/internal/app/device_consumer"
	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
)

func main() {
	mqttClient := infrastructure.NewMqttClient(config.MqttBrokerUri, config.MqttUsername, config.MqttPassword)
	mongo := infrastructure.NewMongoDatabase(config.MongoUri, config.MongoDbName)
	redisClient := infrastructure.NewRedisClient(config.RedisUri)

	repository := device_consumer.NewRepository(mongo)
	consumer := device_consumer.NewConsumer(repository, mqttClient, redisClient)
	consumer.StartConsume()

	select {}
}
