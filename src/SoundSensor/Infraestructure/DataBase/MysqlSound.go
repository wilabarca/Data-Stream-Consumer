package database

import (
	entities "DataConsumer/src/SoundSensor/Domain/Entities"
	repositories "DataConsumer/src/SoundSensor/Domain/Repositories"
	"database/sql"
	"fmt"
	"time"
)

type MySQLSoundRepository struct {
	db *sql.DB
}

func NewSoundRepository(db *sql.DB) repositories.SoundSensor {
	return &MySQLSoundRepository{db: db}
}

func (m *MySQLSoundRepository) SaveSoundData(sensor *entities.SoundSensor) error {
	parsedTime, err := time.Parse(time.RFC3339, sensor.Timestamp)
	if err != nil {
		return fmt.Errorf("error al parsear el timestamp: %v", err)
	}
	mysqlFormattedTime := parsedTime.Format("2006-01-02 15:04:05")

	_, err = m.db.Exec(
		"INSERT INTO SoundSensor (SensorID, RuidoDB, Timestamp) VALUES (?, ?, ?)",
		sensor.SensorID, sensor.RuidoDB, mysqlFormattedTime,
	)
	return err
}

func (m *MySQLSoundRepository) GetSoundData() ([]*entities.SoundSensor, error) {
	rows, err := m.db.Query("SELECT sensorID, ruido_dB, timestamp FROM sound_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*entities.SoundSensor
	for rows.Next() {
		var sensor entities.SoundSensor
		if err := rows.Scan(&sensor.SensorID, &sensor.RuidoDB, &sensor.Timestamp); err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
	}
	return sensors, nil
}