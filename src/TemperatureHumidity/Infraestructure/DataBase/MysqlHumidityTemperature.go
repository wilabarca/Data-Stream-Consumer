package database

import (
	"database/sql"
	entities "DataConsumer/src/TemperatureHumidity/Domain/Entities"
	repositories "DataConsumer/src/TemperatureHumidity/Domain/Repositories"
)

type MySQLTemperatureHumidityRepository struct {
	db *sql.DB
}

func NewTemperatureHumidityRepository(db *sql.DB) repositories.TemperatureHumidityRepository {
	return &MySQLTemperatureHumidityRepository{db: db}
}

func (m *MySQLTemperatureHumidityRepository) SaveTemperatureHumidityData(sensor *entities.TemperatureHumiditySensor) error {
	_, err := m.db.Exec(
		"INSERT INTO temperature_humidity_sensors (sensor_id, temperatura, humedad, timestamp) VALUES (?, ?, ?, ?)",
		sensor.SensorID, sensor.Temperatura, sensor.Humedad, sensor.Timestamp,
	)
	return err
}

func (m *MySQLTemperatureHumidityRepository) GetTemperatureHumidityData() ([]*entities.TemperatureHumiditySensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, temperatura, humedad, timestamp FROM temperature_humidity_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*entities.TemperatureHumiditySensor
	for rows.Next() {
		var sensor entities.TemperatureHumiditySensor
		if err := rows.Scan(&sensor.ID, &sensor.SensorID, &sensor.Temperatura, &sensor.Humedad, &sensor.Timestamp); err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
	}
	return sensors, nil
}