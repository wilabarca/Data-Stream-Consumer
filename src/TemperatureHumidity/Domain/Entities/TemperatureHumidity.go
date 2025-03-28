package entities

type TemperatureHumiditySensor struct {
	ID          int     `json:"id"`
	SensorID    string  `json:"sensor_id"`
	Temperatura float64 `json:"temperatura"`
	Humedad     float64 `json:"humedad"`
	Timestamp   string  `json:"timestamp"`
}