package entities

type SoundSensor struct {
	ID        int    `json:"-"`
	SensorID  string `json:"sensor_id"`
	RuidoDB   int    `json:"nivel"`
	Timestamp string `json:"timestamp"`
}