package helper

import (
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func CommandDeviceWithMqtt(mqttClient mqtt.Client, deviceId string, params map[string]any) {
	topic := fmt.Sprintf("device/%s/command", deviceId)

	paramsJson, err := json.Marshal(params)
	LogIfErr(err)

	token := mqttClient.Publish(topic, 1, false, paramsJson)
	token.Wait()
	if token.Error() != nil {
		log.Println(token.Error())
	}
}
