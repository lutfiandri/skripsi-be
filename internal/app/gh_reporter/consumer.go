package gh_reporter

import (
	"log"

	"skripsi-be/internal/constant"
	"skripsi-be/internal/dto/device_state_log_dto"
	"skripsi-be/internal/util/helper"

	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"google.golang.org/api/homegraph/v1"
)

type Consumer struct {
	redisClient            *redis.Client
	homegraphDeviceService *homegraph.DevicesService
}

func NewConsumer(redisClient *redis.Client, homegraphDeviceService *homegraph.DevicesService) Consumer {
	return Consumer{
		redisClient:            redisClient,
		homegraphDeviceService: homegraphDeviceService,
	}
}

func (consumer *Consumer) StartConsume() {
	stream := constant.RedisDeviceStateStream
	group := constant.RedisGhConsumerGroup

	// Create a consumer group (if it doesn't already exist)
	err := consumer.redisClient.XGroupCreateMkStream(stream, group, "$").Err()
	if err != nil {
		log.Println("Failed to create consumer group: ", err)
	}

	for {
		msgs, err := consumer.redisClient.XReadGroup(&redis.XReadGroupArgs{
			Group:    group,
			Consumer: uuid.NewString(),
			Streams:  []string{stream, ">"},
			Count:    1,
			Block:    0,
		}).Result()
		helper.LogIfErr(err)

		for _, msg := range msgs {
			for _, xmsg := range msg.Messages {
				message := xmsg.Values

				dataField := message[constant.RedisStreamKey].(string)

				data, err := helper.UnmarshalJson[device_state_log_dto.DeviceStateLog[any]]([]byte(dataField))
				helper.LogIfErr(err)

				consumer.HandleIncomingData(data)
			}
		}

		// fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), data)

		// consumer.HandleIncomingData(data)
	}
}

func (consumer *Consumer) HandleIncomingData(data device_state_log_dto.DeviceStateLog[any]) {
	deviceId := data.DeviceId
	state := data.State

	deviceStateMap := map[string]any{}
	deviceStateMap[deviceId] = state

	err := helper.HomegraphReportStateAndNotification(consumer.homegraphDeviceService, data.UserId, deviceStateMap)
	helper.LogIfErr(err)
}
