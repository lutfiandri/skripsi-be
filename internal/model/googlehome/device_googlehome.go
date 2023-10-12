package googlehome

// DeviceName contains different ways of identifying the device
type DeviceName struct {
	// DefaultNames (not user settable)
	DefaultNames []string
	// Name supplied by the user for display purposes
	Name string
	// Nicknames given to this, should a user have multiple ways to refer to the device
	Nicknames []string
}

// DeviceInfo contains different properties of the device
type DeviceInfo struct {
	// Manufacturer of the device
	Manufacturer string
	// Model of the device
	Model string
	// HwVersion of the device
	HwVersion string
	// SwVersion of the device
	SwVersion string
}

// OtherDeviceID contains alternative ways to identify this device.
type OtherDeviceID struct {
	AgentID  string
	DeviceID string
}

// Device represents a single provider-supplied device profile.
type Device struct {
	// ID of the device
	ID string

	// Type of the device.
	// See https://developers.google.com/assistant/smarthome/guides is a list of possible types
	Type string

	// Traits of the device.
	// See https://developers.google.com/assistant/smarthome/traits for a list of possible traits
	// The set of assigned traits will dictate which actions can be performed on the device
	Traits []string

	// Name of the device.
	Name DeviceName

	// WillReportState using the ReportState API (should be true)
	WillReportState bool

	// RoomHint guides Google as to which room this device is in
	RoomHint string

	// Attributes linked to the defined traits
	Attributes map[string]interface{}

	// DeviceInfo that is physically defined
	DeviceInfo DeviceInfo

	// OtherDeviceIDs allows for this to be logically linked to other devices
	OtherDeviceIDs []OtherDeviceID

	// CustomData specified which will be included unmodified in subsequent requests.
	CustomData map[string]interface{}
}
