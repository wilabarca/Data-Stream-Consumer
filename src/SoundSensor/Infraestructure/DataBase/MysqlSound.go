package database

import (
	"database/sql"
	entities "DataConsumer/src/SoundSensor/Domain/Entities"
	repositories "DataConsumer/src/SoundSensor/Domain/Repositories"
)

type MySQLSoundRepository struct {
	db *sql.DB
}

func NewSoundRepository(db *sql.DB) repositories.SoundSensor {
	return &MySQLSoundRepository{db: db}
}

func (m *MySQLSoundRepository) SaveSoundData(sensor *entities.SoundSensor) error {
	_, err := m.db.Exec(
		"INSERT INTO sound_sensors (sensor_id, ruido_dB, timestamp) VALUES (?, ?, ?)",
		sensor.SensorID, sensor.RuidoDB, sensor.Timestamp,
	)
	return err
}

func (m *MySQLSoundRepository) GetSoundData() ([]*entities.SoundSensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, ruido_dB, timestamp FROM sound_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*entities.SoundSensor
	for rows.Next() {
		var sensor entities.SoundSensor
		if err := rows.Scan(&sensor.ID, &sensor.SensorID, &sensor.RuidoDB, &sensor.Timestamp); err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
	}
	return sensors, nil
}