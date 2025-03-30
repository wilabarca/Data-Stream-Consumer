package entities



// SoundSensor representa los datos de un sensor de sonido.
type SoundSensor struct {
	ID          int     `json:"id"`           // ID del sensor en la base de datos
	SensorID    string  `json:"sensor_id"`    // ID del sensor de sonido
	RuidoDB     float64 `json:"ruido_dB"`     // Nivel de ruido en decibelios
	Timestamp   string  `json:"timestamp"`    // Fecha y hora en que se registró la lectura
}
