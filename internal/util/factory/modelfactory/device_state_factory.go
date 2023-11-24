package modelfactory

import (
	"time"

	"skripsi-be/internal/model/domain"
	"skripsi-be/internal/model/kafkamodel"
	"skripsi-be/internal/model/mqttmodel"

	"github.com/google/uuid"
)

// SmartPlug LIT
func DeviceStateSmartPlugLitMqttToKafka(mqttState mqttmodel.SmartPlugLitIncomingData) kafkamodel.DeviceState {
	deviceState := kafkamodel.DeviceState{
		Id:        uuid.NewString(),
		DeviceId:  mqttState.DeviceId,
		CreatedAt: time.Now(),
		State: domain.SmartPlugLitState{
			Volt:         mqttState.Data.Volt,
			MilliAmpere:  mqttState.Data.MilliAmpere,
			Watt:         mqttState.Data.Watt,
			KiloWattHour: mqttState.Data.KiloWattHour,
			On:           mqttState.Data.On,
			NumOfSensor:  mqttState.Data.NumOfSensor,
		},
	}

	return deviceState
}

func DeviceStateSmartPlugLitMqttToDomain(mqttState mqttmodel.SmartPlugLitIncomingData) domain.DeviceState {
	deviceState := domain.DeviceState{
		Id:        uuid.NewString(),
		DeviceId:  mqttState.DeviceId,
		CreatedAt: time.Now(),
		State: domain.SmartPlugLitState{
			Volt:         mqttState.Data.Volt,
			MilliAmpere:  mqttState.Data.MilliAmpere,
			Watt:         mqttState.Data.Watt,
			KiloWattHour: mqttState.Data.KiloWattHour,
			On:           mqttState.Data.On,
			NumOfSensor:  mqttState.Data.NumOfSensor,
		},
	}

	return deviceState
}
