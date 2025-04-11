package entities

type LightSensor struct {
	ID        int     `json:"id"`
	SensorID  string  `json:"sensor_id"`
	Nivel     float64 `json:"nivel"`
	Timestamp string  `json:"timestamp"`
}