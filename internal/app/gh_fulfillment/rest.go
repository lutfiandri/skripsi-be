package gh_fulfillment

import "skripsi-be/internal/util/gh_builder"

// https://developers.home.google.com/cloud-to-cloud/integration/query-execute#execute-response

// request
type DeviceRequest struct {
	Id         string         `json:"id"`
	CustomData map[string]any `json:"customData"`
}

type ExecutionRequest struct {
	Command string `json:"command"`
	Params  map[string]any
}

type CommandRequest struct {
	Devices   []DeviceRequest    `json:"devices"`
	Execution []ExecutionRequest `json:"execution"`
}

type Request struct {
	RequestId string `json:"requestId"`
	Inputs    []struct {
		Intent  string `json:"intent"`
		Payload struct {
			Devices  []DeviceRequest  `json:"devices"`
			Commands []CommandRequest `json:"commands"`
		} `json:"payload"`
	} `json:"inputs"`
}

// SYNC response
type SyncPayloadResponse struct {
	AgentUserId string              `json:"agentUserId"`
	Devices     []gh_builder.Device `json:"devices"`
}

type SyncResponse struct {
	RequestId string              `json:"requestId"`
	Payload   SyncPayloadResponse `json:"payload"`
}

// QUERY response
type QueryPayloadResponse struct {
	Devices map[string]any `json:"devices"`
}

type QueryResponse struct {
	RequestId string               `json:"requestId"`
	Payload   QueryPayloadResponse `json:"payload"`
}

// EXECUTE response
type ExecuteCommandResponse struct {
	Ids       []string       `json:"ids"`
	Status    string         `json:"status"`
	States    map[string]any `json:"states"`
	ErrorCode string         `json:"errorCode"`
}

type ExecutePayloadResponse struct {
	Commands []ExecuteCommandResponse `json:"commands"`
}

type ExecuteResponse struct {
	RequestId string                 `json:"requestId"`
	Payload   ExecutePayloadResponse `json:"payload"`
}

// DISCONNECT response
type DisconnectResponse struct{}
