package database

import (
	entities "DataConsumer/src/TemperatureHumidity/Domain/Entities"
	repositories "DataConsumer/src/TemperatureHumidity/Domain/Repositories"
	"database/sql"
	"fmt"
	"time"
)

type MySQLTemperatureHumidityRepository struct {
	db *sql.DB
}

func NewTemperatureHumidityRepository(db *sql.DB) repositories.TemperatureHumidityRepository {
	return &MySQLTemperatureHumidityRepository{db: db}
}

func (m *MySQLTemperatureHumidityRepository) SaveTemperatureHumidityData(sensor *entities.TemperatureHumiditySensor) error {
	parsedTime, err := time.Parse(time.RFC3339, sensor.Timestamp)
	if err != nil {
		return fmt.Errorf("error al parsear el timestamp: %v", err)
	}
	mysqlFormattedTime := parsedTime.Format("2006-01-02 15:04:05")

	_, err = m.db.Exec(
		"INSERT INTO TemperatureHumidity (SensorID, Temperature, Humidity, Timestamp) VALUES (?, ?, ?, ?)",
		sensor.SensorID, sensor.Temperature, sensor.Humidity, mysqlFormattedTime,
	)
	return err
}
func (m *MySQLTemperatureHumidityRepository) GetTemperatureHumidityData() ([]*entities.TemperatureHumiditySensor, error) {
	rows, err := m.db.Query("SELECT sensorID, temperature, humidity, timestamp FROM temperature_humidity")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*entities.TemperatureHumiditySensor
	for rows.Next() {
		var sensor entities.TemperatureHumiditySensor
		if err := rows.Scan(&sensor.SensorID, &sensor.Temperature, &sensor.Humidity, &sensor.Timestamp); err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
	}
	return sensors, nil
}
