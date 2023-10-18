package main

import (
	"log"
	"sync"

	"skripsi-be/internal/consumer"
	"skripsi-be/internal/infrastructure"
)

func main() {
	log.Println("device consumer service")
	mqttClient := infrastructure.NewMqttClient("tcp://broker.emqx.io:1883")

	deviceConsumer := consumer.NewDeviceConsumer(mqttClient)
	deviceConsumer.InitMqttSubscriber()

	// prevent app from exiting
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
