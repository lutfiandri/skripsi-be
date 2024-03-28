package gh_builder

func NewLightBuilder() *BaseDeviceBuilder {
	builder := NewBaseDeviceBuilder().
		SetType("action.devices.types.LIGHT").
		AddTraits("action.devices.traits.OnOff")
	return builder
}
