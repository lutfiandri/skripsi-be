package gh_builder

type Device struct {
	ID              string         `json:"id"`
	Type            string         `json:"type"`
	WillReportState bool           `json:"willReportState"`
	Traits          []string       `json:"traits"`
	Attributes      map[string]any `json:"attributes"`
	Name            struct {
		DefaultNames []string `json:"defaultNames"`
		Name         string   `json:"name"`
		Nicknames    []string `json:"nicknames"`
	} `json:"name"`
	RoomHint   string `json:"roomHint"`
	DeviceInfo struct {
		Manufacturer string `json:"manufacturer"`
		Model        string `json:"model"`
		HWVersion    string `json:"hwVersion"`
		SWVersion    string `json:"swVersion"`
	} `json:"deviceInfo"`
	CustomData map[string]any `json:"customData"`
}

// BaseDeviceBuilder is the base builder for all device types
type BaseDeviceBuilder struct {
	device *Device // Internal device instance
}

// NewBaseDeviceBuilder creates a new BaseDeviceBuilder instance
func NewBaseDeviceBuilder() *BaseDeviceBuilder {
	return &BaseDeviceBuilder{device: &Device{}}
}

// SetID sets the unique identifier for the device
func (b *BaseDeviceBuilder) SetID(value string) *BaseDeviceBuilder {
	b.device.ID = value
	return b
}

// SetType sets the device's type (e.g., action.devices.types.LIGHT)
func (b *BaseDeviceBuilder) SetType(value string) *BaseDeviceBuilder {
	b.device.Type = value
	return b
}

// SetWillReportState indicates whether the device will automatically report its state changes
func (b *BaseDeviceBuilder) SetWillReportState(value bool) *BaseDeviceBuilder {
	b.device.WillReportState = value
	return b
}

// AddTraits adds one or more traits to the device
func (b *BaseDeviceBuilder) AddTraits(values ...string) *BaseDeviceBuilder {
	b.device.Traits = append(b.device.Traits, values...)
	return b
}

// SetAttributes sets the device's attributes (current state)
func (b *BaseDeviceBuilder) SetAttributes(attributes map[string]any) *BaseDeviceBuilder {
	b.device.Attributes = attributes
	return b
}

// SetName defines the device's name, default names, and nicknames
func (b *BaseDeviceBuilder) SetName(defaultNames []string, name string, nicknames []string) *BaseDeviceBuilder {
	b.device.Name.DefaultNames = defaultNames
	b.device.Name.Name = name
	b.device.Name.Nicknames = nicknames
	return b
}

// SetRoomHint specifies the room where the device is located (optional)
func (b *BaseDeviceBuilder) SetRoomHint(value string) *BaseDeviceBuilder {
	b.device.RoomHint = value
	return b
}

// SetDeviceInfo provides details about the device's manufacturer, model, etc.
func (b *BaseDeviceBuilder) SetDeviceInfo(manufacturer string, model string, hwVersion string, swVersion string) *BaseDeviceBuilder {
	b.device.DeviceInfo.Manufacturer = manufacturer
	b.device.DeviceInfo.Model = model
	b.device.DeviceInfo.HWVersion = hwVersion
	b.device.DeviceInfo.SWVersion = swVersion
	return b
}

// SetCustomData stores application-specific data associated with the device (optional)
func (b *BaseDeviceBuilder) SetCustomData(data map[string]any) *BaseDeviceBuilder {
	b.device.CustomData = data
	return b
}

// Build builds and returns the final Device object
func (b *BaseDeviceBuilder) Build() Device {
	return *b.device
}
