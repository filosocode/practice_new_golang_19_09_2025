package device

import (
	"math/rand"

	"time"
)

// ConsumptionData representa un dato de consumo electrico
// estructura con deviceID, timestamp y valor en kWh.
type ConsumptionData struct {
	DeviceID  string
	Timestamp time.Time
	Value     float64
}

// SimulateConsuption genera un dato simulado de consumo
func SimulateConsumption(DeviceID string) ConsumptionData {
	return ConsumptionData{
		DeviceID:  DeviceID,
		Timestamp: time.Now(),
		Value:     0.5 + rand.Float64()*1.5, //entre 0.5 y 2.0 kwh
	}
}
