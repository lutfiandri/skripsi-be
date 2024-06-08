package helper

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"google.golang.org/api/homegraph/v1"
)

func HomegraphRequestSync(homegraphDeviceService *homegraph.DevicesService, agentUserId string) error {
	request := homegraph.RequestSyncDevicesRequest{
		AgentUserId: agentUserId,
		Async:       false,
	}

	log.Println("[HOMEGRAPH] Request Sync for Agent User ID:", agentUserId)

	_, err := homegraphDeviceService.RequestSync(&request).Do()
	return err
}

func HomegraphReportStateAndNotification(homegraphDeviceService *homegraph.DevicesService, agentUserId string, states map[string]any) error {
	raw, err := json.Marshal(states)
	if err != nil {
		return err
	}

	request := homegraph.ReportStateAndNotificationRequest{
		AgentUserId: agentUserId,
		RequestId:   uuid.NewString(),
		Payload: &homegraph.StateAndNotificationPayload{
			Devices: &homegraph.ReportStateAndNotificationDevice{
				States: raw,
			},
		},
	}

	_, err = homegraphDeviceService.ReportStateAndNotification(&request).Do()
	return err
}
