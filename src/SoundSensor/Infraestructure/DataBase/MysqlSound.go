package database

import (
	entities "DataConsumer/src/SoundSensor/Domain/Entities"
	repositories "DataConsumer/src/SoundSensor/Domain/Repositories"
	"database/sql"
)

type MySQLSoundSensorRepository struct {
	db *sql.DB
}

// Constructor para el repository
func NewMySQLSoundSensorRepository(db *sql.DB) repositories.SoundSensor {
	return &MySQLSoundSensorRepository{db: db}
}

// SaveSoundData guarda los datos del sensor de sonido en la base de datos.
func (m *MySQLSoundSensorRepository) SaveSoundData(sensor *entities.SoundSensor) error {
	_, err := m.db.Exec("INSERT INTO sound_sensors (sensor_id, ruido_dB, timestamp) VALUES (?, ?, ?)",
		sensor.SensorID, sensor.RuidoDB, sensor.Timestamp)
	return err
}

// GetSoundData obtiene todos los datos de los sensores de sonido desde la base de datos.
func (m *MySQLSoundSensorRepository) GetSoundData() ([]*entities.SoundSensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, ruido_dB, timestamp FROM sound_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var soundSensors []*entities.SoundSensor
	for rows.Next() {
		var sensor entities.SoundSensor
		if err := rows.Scan(&sensor.ID, &sensor.SensorID, &sensor.RuidoDB, &sensor.Timestamp); err != nil {
			return nil, err
		}
		soundSensors = append(soundSensors, &sensor)
	}
	return soundSensors, nil
}
