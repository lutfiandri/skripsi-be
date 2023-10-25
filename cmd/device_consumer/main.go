package main

import (
	"log"
	"sync"

	"skripsi-be/internal/config"
	"skripsi-be/internal/consumer"
	"skripsi-be/internal/infrastructure"
	"skripsi-be/internal/publisher"
)

func main() {
	log.Println("device consumer service")
	mqttClient := infrastructure.NewMqttClient(config.MqttBrokerUri, config.MqttUsername, config.MqttPassword)
	kafkaWriter := infrastructure.NewKafkaWriter("localhost:9092")

	kafkaPublisher := publisher.NewKafkaPublisher(kafkaWriter)

	deviceConsumer := consumer.NewDeviceConsumer(mqttClient, kafkaPublisher)
	deviceConsumer.InitMqttSubscriber()

	// prevent app from exiting
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
