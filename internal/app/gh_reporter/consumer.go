package gh_reporter

import (
	"context"
	"encoding/json"

	"skripsi-be/internal/dto/device_state_log_dto"
	"skripsi-be/internal/util/helper"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"google.golang.org/api/homegraph/v1"
)

type Consumer struct {
	kafkaReader            *kafka.Reader
	homegraphDeviceService *homegraph.DevicesService
}

func NewConsumer(kafkaReader *kafka.Reader) Consumer {
	homegraphService, err := homegraph.NewService(context.Background())
	helper.PanicIfErr(err)

	homegraphDeviceService := homegraph.NewDevicesService(homegraphService)

	return Consumer{
		kafkaReader:            kafkaReader,
		homegraphDeviceService: homegraphDeviceService,
	}
}

func (consumer *Consumer) StartConsume() {
	for {
		m, err := consumer.kafkaReader.ReadMessage(context.Background())
		helper.PanicIfErr(err)

		data := device_state_log_dto.DeviceStateLog[any]{}
		err = json.Unmarshal(m.Value, &data)
		helper.LogIfErr(err)

		// fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), data)

		consumer.HandleIncomingData(data)
	}
}

func (consumer *Consumer) HandleIncomingData(data device_state_log_dto.DeviceStateLog[any]) {
	deviceId := data.DeviceId
	state := data.State

	deviceStateMap := map[string]any{}
	deviceStateMap[deviceId] = state

	raw, err := json.Marshal(deviceStateMap)
	helper.LogIfErr(err)

	request := homegraph.ReportStateAndNotificationRequest{
		AgentUserId: data.UserId,
		RequestId:   uuid.NewString(),
		Payload: &homegraph.StateAndNotificationPayload{
			Devices: &homegraph.ReportStateAndNotificationDevice{
				States: raw,
			},
		},
	}

	_, err = consumer.homegraphDeviceService.ReportStateAndNotification(&request).Do()
	helper.LogIfErr(err)

	// dataJson, err := request.MarshalJSON()
	// helper.LogIfErr(err)

	// log.Printf("%+v\n\n", string(dataJson))
}
