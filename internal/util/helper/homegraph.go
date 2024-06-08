package helper

import (
	"log"

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
