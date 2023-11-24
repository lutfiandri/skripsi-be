package consumer

import (
	"context"
	"fmt"
	"log"

	"skripsi-be/internal/model/mqttmodel"
	"skripsi-be/internal/publisher"
	"skripsi-be/internal/repository"
	"skripsi-be/internal/util/factory/modelfactory"
	"skripsi-be/internal/util/helper"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type DeviceConsumer interface {
	InitMqttSubscriber()

	HandleSmartPlugData(client mqtt.Client, message mqtt.Message)
	HandleSmartLampData(client mqtt.Client, message mqtt.Message)
}

type deviceConsumer struct {
	mqttClient            mqtt.Client
	kafkaPublisher        publisher.KafkaPublisher
	deviceStateRepository repository.DeviceStateRepository
}

func NewDeviceConsumer(mqttClient mqtt.Client, kafkaPublisher publisher.KafkaPublisher, deviceStateRepository repository.DeviceStateRepository) DeviceConsumer {
	return &deviceConsumer{
		mqttClient:            mqttClient,
		kafkaPublisher:        kafkaPublisher,
		deviceStateRepository: deviceStateRepository,
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
			topic:    "+/meter/data",
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
	data, err := helper.UnmarshalJson[mqttmodel.SmartPlugLitIncomingData](message.Payload())
	if err != nil {
		log.Println("error on parsing data: ", err.Error())
		return
	}
	log.Printf("smartplug lit: %+v\n", data)

	deviceState := modelfactory.DeviceStateSmartPlugLitMqttToDomain(data)
	if err := consumer.deviceStateRepository.InsertDeviceState(context.Background(), deviceState); err != nil {
		log.Println("error on saving SmartPlugLit state to database: ", err.Error())
	}

	kafkaMessage := modelfactory.DeviceStateSmartPlugLitMqttToKafka(data)

	if err := consumer.kafkaPublisher.Publish(context.Background(), "test-topic", kafkaMessage); err != nil {
		log.Println("error on publishing data: ", err.Error())
	} else {
		log.Println("sent to kafka: ", kafkaMessage)
	}
}

func (consumer *deviceConsumer) HandleSmartLampData(client mqtt.Client, message mqtt.Message) {
	panic("unimplemented")
}
