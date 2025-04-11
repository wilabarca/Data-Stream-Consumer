package entities

type AirQualitySensor struct {
	ID          int     `json:"id"`
	SensorID    string  `json:"sensor_id"`
	CO2PPM      int     `json:"co2_ppm"`
	Air_level        int     `json:"air_level"`
	Timestamp   string  `json:"timestamp"`
}
