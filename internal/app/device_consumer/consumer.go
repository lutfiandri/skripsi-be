package device_consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/domain"
	"skripsi-be/internal/dto/device_state_log_dto"
	"skripsi-be/internal/util/helper"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
)

type Consumer interface {
	StartConsume()
	HandleIncomingData(client mqtt.Client, message mqtt.Message)
	GetRedisKey(dto device_state_log_dto.DeviceStateLog[any]) string
}

type consumer struct {
	repository  Repository
	mqttClient  mqtt.Client
	redisClient *redis.Client
}

func NewConsumer(repository Repository, mqttClient mqtt.Client, redisClient *redis.Client) Consumer {
	return &consumer{
		repository:  repository,
		mqttClient:  mqttClient,
		redisClient: redisClient,
	}
}

func (consumer consumer) StartConsume() {
	// topic: /device/{device_id}/state

	go func() {
		topic := "device/+/state"
		token := consumer.mqttClient.Subscribe(topic, 1, consumer.HandleIncomingData)
		if token.Wait() && token.Error() != nil {
			log.Println(fmt.Sprintf("error on subscribing %s:", topic), token.Error())
		} else {
			log.Printf("subscribing %s\n", topic)
		}
	}()
}

func (consumer consumer) HandleIncomingData(client mqtt.Client, message mqtt.Message) {
	log.Println("mqtt data:", string(message.Payload()))
	data_dto, err := helper.UnmarshalJson[device_state_log_dto.DeviceStateLog[any]](message.Payload())
	if err != nil {
		log.Println(err.Error())
		return
	}

	deviceId, err := uuid.Parse(data_dto.DeviceId)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// insert to db
	data_domain := domain.DeviceStateLog[any]{
		Id:        uuid.New(),
		DeviceId:  deviceId,
		State:     data_dto.State,
		CreatedAt: data_dto.CreatedAt,
	}

	err = consumer.repository.InsertDeviceState(context.TODO(), data_domain)
	helper.LogIfErr(err)

	err = consumer.repository.UpdateDeviceLastState(context.TODO(), data_domain)
	helper.LogIfErr(err)

	// get user id
	device, err := consumer.repository.GetDeviceById(context.TODO(), deviceId)
	if err != nil {
		log.Println("GetDeviceById:", err.Error())
		return
	}

	if device.UserId == nil {
		log.Println("No User ID")
		return
	}

	data_dto.UserId = device.UserId.String()

	data_map := helper.StructToMap(data_dto)
	log.Println(data_map)

	consumer.PublishToRedisStream(data_dto)

	// consumer.PublishToKafka(data_domain.Id.String(), data_dto)

	// PIPELINE TO KAFKA
	// 1. Get from redis
	// 2. If exists && same state -> don't write to kafka. If not -> write to kafka
	// 3. Set to redis with expiration (ex: 1 minutes)
	// note: redis value is stringify of state

	// // 1. Get from redis
	redis_key := consumer.GetRedisKey(data_dto)
	// value, err := consumer.redisClient.Get(redis_key).Result()

	// // 2. If exists && same state -> don't write to kafka. If not -> write to kafka
	// if err != nil {
	// 	log.Println(err)
	// 	consumer.PublishToKafka(data_domain.Id.String(), data_dto)
	// } else {
	// 	data_redis, err := helper.UnmarshalJson[any]([]byte(value))
	// 	log.Println(data_redis)
	// 	if err != nil {
	// 		log.Println(err)
	// 		consumer.PublishToKafka(data_domain.Id.String(), data_dto)
	// 	} else {
	// 		// check same
	// 		same := reflect.DeepEqual(data_redis, data_dto.State)
	// 		if !same {
	// 			consumer.PublishToKafka(data_domain.Id.String(), data_dto)
	// 		}
	// 	}
	// }

	// 3. Set to redis. Expires in 1 minute
	data_dto_bytes, err := json.Marshal(data_dto.State)
	helper.LogIfErr(err)
	// FIXME: set to one minute
	err = consumer.redisClient.Set(redis_key, string(data_dto_bytes), 16*time.Second).Err()
	helper.LogIfErr(err)
}

func (consumer consumer) GetRedisKey(dto device_state_log_dto.DeviceStateLog[any]) string {
	key := fmt.Sprintf("device:%s", dto.DeviceId)
	return key
}

// func (consumer consumer) PublishToKafka(key string, data_dto device_state_log_dto.DeviceStateLog[any]) {
// 	// insert to kafka
// 	data_kafka, err := json.Marshal(data_dto)
// 	helper.LogIfErr(err)
// 	kafka_topic := config.KafkaTopicDeviceState

// 	err = consumer.kafkaWriter.WriteMessages(context.TODO(), kafka.Message{
// 		Topic: kafka_topic,
// 		Key:   []byte(key),
// 		Value: data_kafka,
// 	})
// 	helper.LogIfErr(err)
// }

func (consumer *consumer) PublishToRedisStream(data_dto device_state_log_dto.DeviceStateLog[any]) {
	stream := constant.RedisDeviceStateStream

	jsonData, err := json.Marshal(data_dto)
	if err != nil {
		log.Println(err)
		return
	}

	// redis map can only 1 depth
	var message map[string]any = map[string]any{
		constant.RedisStreamKey: jsonData,
	}

	err = consumer.redisClient.XAdd(&redis.XAddArgs{
		// ID:     uuid.NewString(),
		Stream: stream,
		Values: message,
	}).Err()

	helper.LogIfErr(err)
}
