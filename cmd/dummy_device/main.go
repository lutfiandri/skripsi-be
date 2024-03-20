package main

import (
	"sync"
	"time"

	"skripsi-be/internal/config"
	"skripsi-be/internal/infrastructure"
	devicestatedto "skripsi-be/internal/model/device_state_dto"
)

func main() {
	mqttClient := infrastructure.NewMqttClient(config.MqttBrokerUri, config.MqttPassword, config.MqttPassword)

	go func() {
		lightState := devicestatedto.LightState{}

		time.Sleep(10 * time.Second)
	}()

	// prevent app from exiting
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
