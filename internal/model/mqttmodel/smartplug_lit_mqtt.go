package mqttmodel

import (
	"strings"
	"time"
)

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	dateFormat := "2006-01-02 15:04:05"
	timeString := strings.Trim(string(b), `"`)
	t, err := time.Parse(dateFormat, timeString)
	if err != nil {
		return err
	}

	*ct = CustomTime(t)
	return nil
}

type SmartPlugLitIncomingData struct {
	DeviceId  string     `json:"DevID"`
	Timestamp CustomTime `json:"DateTime"`
	Data      struct {
		Volt         float64 `json:"V"`
		MilliAmpere  float64 `json:"mA"`
		Watt         float64 `json:"W"`
		KiloWattHour float64 `json:"kWh"`
		On           bool    `json:"relay"`
		NumOfSensor  int     `json:"Sensor"`
	} `json:"data"`
}
