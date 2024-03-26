package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	lightDeviceId := "8cd87119-c57b-4a69-be6f-f837adb080eb"
	go loopLight(client, 5*time.Second, lightDeviceId)

	// prevent app from exiting
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func loopLight(client mqtt.Client, interval time.Duration, deviceId string) {
	deviceTypeId := "dfb08e53-2a45-4b63-aa2d-e0795929cffe"
	for {
		topic := fmt.Sprintf("device/%s/%s/state", deviceTypeId, deviceId)

		data := device_state_log_dto.DeviceStateLog[device_state_log_dto.LightState]{
			DeviceId: deviceId,
			State: device_state_log_dto.LightState{
				On: true,
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
