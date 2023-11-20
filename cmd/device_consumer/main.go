package main

import (
	"log"
	"sync"

	"skripsi-be/internal/config"
	"skripsi-be/internal/consumer"
	"skripsi-be/internal/infrastructure"
	"skripsi-be/internal/publisher"
	"skripsi-be/internal/repository"
)

func main() {
	log.Println("device consumer service")
	mongo := infrastructure.NewMongoDatabase(config.MongoUri, config.MongoDbName)

	mqttClient := infrastructure.NewMqttClient(config.MqttBrokerUri, config.MqttUsername, config.MqttPassword)
	kafkaWriter := infrastructure.NewKafkaWriter("localhost:9092")

	kafkaPublisher := publisher.NewKafkaPublisher(kafkaWriter)

	deviceStateRepository := repository.NewDeviceStateRepository(mongo, "device_states")

	deviceConsumer := consumer.NewDeviceConsumer(mqttClient, kafkaPublisher, deviceStateRepository)
	deviceConsumer.InitMqttSubscriber()

	// prevent app from exiting
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
