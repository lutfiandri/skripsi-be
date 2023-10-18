package consumer

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DeviceConsumer interface {
	InitMqttSubscriber()

	HandleSmartPlugData(client mqtt.Client, message mqtt.Message)
	HandleSmartLampData(client mqtt.Client, message mqtt.Message)
}

type deviceConsumer struct {
	mqttClient mqtt.Client
}

func NewDeviceConsumer(mqttClient mqtt.Client) DeviceConsumer {
	return &deviceConsumer{
		mqttClient: mqttClient,
	}
}

func (consumer *deviceConsumer) InitMqttSubscriber() {
	type topicType struct {
		topic    string
		qos      byte
		callback mqtt.MessageHandler
	}
	topics := []topicType{
		{
			topic:    "dteti/capstone-c12",
			qos:      1,
			callback: consumer.HandleSmartPlugData,
		},
	}

	for _, topic := range topics {
		go func(topic topicType) {
			if token := consumer.mqttClient.Subscribe(topic.topic, topic.qos, topic.callback); token.Wait() && token.Error() != nil {
				log.Println(fmt.Sprintf("error on subscribing %s:", topic.topic), token.Error())
			} else {
				log.Printf("subscribing %s\n", topic.topic)
			}
		}(topic)
	}
}

func (consumer *deviceConsumer) HandleSmartPlugData(client mqtt.Client, message mqtt.Message) {
	log.Println("message", message.Topic(), string(message.Payload()))
}

func (consumer *deviceConsumer) HandleSmartLampData(client mqtt.Client, message mqtt.Message) {
	panic("unimplemented")
}
