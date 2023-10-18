package infrastructure

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMqttClient(brokerUri string) mqtt.Client {
	opts := mqtt.NewClientOptions().AddBroker(brokerUri)

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}

	opts.SetDefaultPublishHandler(messagePubHandler)

	client := mqtt.NewClient(opts)

	// ping
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("error on connecting mqtt: ", token.Error())
	} else {
		log.Printf("mqtt client connected to %s\n", brokerUri)
	}

	return client
}
