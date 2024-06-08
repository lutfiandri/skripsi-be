package helper

import "google.golang.org/api/homegraph/v1"

func HomegraphRequestSync(homegraphDeviceService *homegraph.DevicesService, agentUserId string) error {
	request := homegraph.RequestSyncDevicesRequest{
		AgentUserId: agentUserId,
		Async:       false,
	}

	_, err := homegraphDeviceService.RequestSync(&request).Do()
	return err
}
