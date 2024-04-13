package device_state_log_dto

import (
	"time"
)

type DeviceStateLog[T any] struct {
	Id        string    `json:"id"`
	DeviceId  string    `json:"device_id"`
	State     T         `json:"state"`
	CreatedAt time.Time `json:"created_at"`
	UserId    string    `json:"user_id"` // used from consumer to kafka
}
