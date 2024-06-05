package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"skripsi-be/internal/config"
	"skripsi-be/internal/dto/device_state_log_dto"
	"skripsi-be/internal/infrastructure"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	log.Println("dummy device service")

	client := infrastructure.NewMqttClient(config.MqttBrokerUri, config.MqttUsername, config.MqttPassword)

	lightDeviceId := "e977c236-48f8-43cd-b987-e0b55b58548c"
	go loopLight(client, 5*time.Second, lightDeviceId)

	// prevent app from exiting
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func loopLight(client mqtt.Client, interval time.Duration, deviceId string) {
	for {
		topic := fmt.Sprintf("device/%s/state", deviceId)

		on := false
		randOn := rand.Intn(2)
		if randOn == 1 {
			on = true
		}

		data := device_state_log_dto.DeviceStateLog[device_state_log_dto.LightState]{
			DeviceId: deviceId,
			State: device_state_log_dto.LightState{
				On: on,
			},
			CreatedAt: time.Now(),
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling data:", err)
			return
		}

		token := client.Publish(topic, 2, true, string(jsonData))
		token.Wait()
		err = token.Error()
		if err != nil {
			fmt.Printf("Light Error: %s\n", err.Error())
		} else {
			fmt.Printf("Light Success: %s\n", string(jsonData))
		}

		time.Sleep(interval)

	}
}
