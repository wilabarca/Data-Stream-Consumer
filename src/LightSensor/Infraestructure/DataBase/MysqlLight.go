package database

import (
	entities "DataConsumer/src/LightSensor/Domain/Entities"
	repositories "DataConsumer/src/LightSensor/Domain/Repositories"
	"database/sql"
	"fmt"
	"time"
)

type MySQLLightRepository struct {
	db *sql.DB
}

func NewLightRepository(db *sql.DB) repositories.LightRepository {
	return &MySQLLightRepository{db: db}
}

func (m *MySQLLightRepository) SaveLightData(sensor *entities.LightSensor) error {
	parsedTime, err := time.Parse(time.RFC3339, sensor.Timestamp)
	if err != nil {
		return fmt.Errorf("error al parsear el timestamp: %v", err)
	}
	mysqlFormattedTime := parsedTime.Format("2006-01-02 15:04:05")

	_, err = m.db.Exec(
		"INSERT INTO LightSensor (SensorID, nivel, Timestamp) VALUES (?, ?, ?)",
		sensor.SensorID, sensor.Nivel, mysqlFormattedTime,
	)
	if err != nil {
		return fmt.Errorf("error al insertar datos del sensor de luz: %v", err)
	}
	return nil
}

func (m *MySQLLightRepository) GetLightData() ([]*entities.LightSensor, error) {
	rows, err := m.db.Query("SELECT sensorID, nivel, timestamp FROM light_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lights []*entities.LightSensor
	for rows.Next() {
		var light entities.LightSensor
		if err := rows.Scan(&light.SensorID, &light.Nivel, &light.Timestamp); err != nil {
			return nil, err
		}
		lights = append(lights, &light)
	}
	return lights, nil
}
