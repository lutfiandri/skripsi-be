package device_consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"skripsi-be/internal/domain"
	"skripsi-be/internal/dto/device_state_log_dto"
	"skripsi-be/internal/util/helper"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Consumer interface {
	StartConsume()
	HandleIncomingData(client mqtt.Client, message mqtt.Message)
}

type consumer struct {
	repository  Repository
	mqttClient  mqtt.Client
	kafkaWriter *kafka.Writer
}

func NewConsumer(repository Repository, mqttClient mqtt.Client, kafkaWriter *kafka.Writer) Consumer {
	return &consumer{
		repository:  repository,
		mqttClient:  mqttClient,
		kafkaWriter: kafkaWriter,
	}
}

func (consumer consumer) StartConsume() {
	// topic: /device/{device_type_id}/{device_id}/state

	go func() {
		topic := "device/+/+/state"
		token := consumer.mqttClient.Subscribe(topic, 1, consumer.HandleIncomingData)
		if token.Wait() && token.Error() != nil {
			log.Println(fmt.Sprintf("error on subscribing %s:", topic), token.Error())
		} else {
			log.Printf("subscribing %s\n", topic)
		}
	}()
}

func (consumer consumer) HandleIncomingData(client mqtt.Client, message mqtt.Message) {
	data_dto, err := helper.UnmarshalJson[device_state_log_dto.DeviceStateLog[any]](message.Payload())
	helper.LogIfErr(err)
	log.Println(data_dto)

	deviceId, err := uuid.Parse(data_dto.DeviceId)
	helper.LogIfErr(err)

	// insert to db
	data_domain := domain.DeviceStateLog[any]{
		Id:        uuid.New(),
		DeviceId:  deviceId,
		State:     data_dto.State,
		CreatedAt: data_dto.CreatedAt,
	}

	err = consumer.repository.InsertDeviceState(context.TODO(), data_domain)
	helper.LogIfErr(err)

	// insert to kafka
	data_kafka, err := json.Marshal(data_dto)
	helper.LogIfErr(err)
	kafka_topic := "device_state"

	err = consumer.kafkaWriter.WriteMessages(context.TODO(), kafka.Message{
		Topic: kafka_topic,
		Key:   []byte(data_domain.Id.String()),
		Value: data_kafka,
	})
	helper.LogIfErr(err)
}
