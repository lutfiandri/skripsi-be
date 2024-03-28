package device_type

import "time"

type GoogleHome struct {
	Type            string   `json:"type" validate:"required"`
	Traits          []string `json:"traits" validate:"required"`
	WillReportState bool     `json:"will_report_state" validate:"required"`
}

type DeviceTypeResponse struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	GoogleHome  GoogleHome `json:"google_home"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type GetDeviceTypeRequest struct {
	Id string `params:"id"`
}

type CreateDeviceTypeRequest struct {
	Name        string     `json:"name" validate:"required"`
	GoogleHome  GoogleHome `json:"google_home" validate:"required"`
	Description string     `json:"description" validate:"required"`
}

type UpdateDeviceTypeRequest struct {
	Id          string     `params:"id"`
	Name        string     `json:"name" validate:"required"`
	GoogleHome  GoogleHome `json:"google_home" validate:"required"`
	Description string     `json:"description" validate:"required"`
}

type DeleteDeviceTypeRequest struct {
	Id string `params:"id"`
}
