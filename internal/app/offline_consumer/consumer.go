package offline_consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/dto/device_state_log_dto"
	"skripsi-be/internal/util/helper"

	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
)

type Consumer interface {
	StartConsume()
}

type consumer struct {
	repository  Repository
	redisClient *redis.Client
}

func NewConsumer(repository Repository, redisClient *redis.Client) Consumer {
	return &consumer{
		repository:  repository,
		redisClient: redisClient,
	}
}

func (consumer consumer) StartConsume() {
	pubsub := consumer.redisClient.PSubscribe("__keyevent@0__:expired")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			fmt.Printf("Error receiving message: %v\n", err)
			return
		}
		fmt.Printf("Received expired key event: %s\n", msg.Payload)

		// example msg.Payload: device:a7371310-85b8-483e-ae84-edb5d3fbe39b
		parts := strings.Split(msg.Payload, ":")

		switch parts[0] {
		case "device":
			deviceId := uuid.MustParse(parts[1])
			err := consumer.repository.UpdateDeviceOnline(context.Background(), deviceId, false)
			helper.LogIfErr(err)

			device, err := consumer.repository.GetDeviceById(context.Background(), deviceId)
			helper.LogIfErr(err)

			data_dto := device_state_log_dto.DeviceStateLog[any]{
				DeviceId: deviceId.String(),
				State:    device.LastState,
				UserId:   device.UserId.String(),
			}

			consumer.PublishToRedisStream(data_dto)
		}
	}
}

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
