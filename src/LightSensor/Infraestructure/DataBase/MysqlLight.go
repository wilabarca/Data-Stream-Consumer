package database

import (
	"database/sql"
	entities "DataConsumer/src/LightSensor/Domain/Entities"
	repositories "DataConsumer/src/LightSensor/Domain/Repositories"
)

type MySQLLightRepository struct {
	db *sql.DB
}

func NewLightRepository(db *sql.DB) repositories.LightRepository {
	return &MySQLLightRepository{db: db}
}

func (m *MySQLLightRepository) SaveLightData(light *entities.LightSensor) error {
	_, err := m.db.Exec(
		"INSERT INTO light_sensors (sensor_id, intensidad, color, estado, timestamp) VALUES (?, ?, ?, ?, ?)",
		light.SensorID, light.Intensidad, light.Color, light.Estado, light.Timestamp,
	)
	return err
}

func (m *MySQLLightRepository) GetLightData() ([]*entities.LightSensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, intensidad, color, estado, timestamp FROM light_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lights []*entities.LightSensor
	for rows.Next() {
		var light entities.LightSensor
		if err := rows.Scan(&light.ID, &light.SensorID, &light.Intensidad, &light.Color, &light.Estado, &light.Timestamp); err != nil {
			return nil, err
		}
		lights = append(lights, &light)
	}
	return lights, nil
}

func (m *MySQLLightRepository) GetLightDataBySensorID(sensorID string) ([]*entities.LightSensor, error) {
	rows, err := m.db.Query(
		"SELECT id, sensor_id, intensidad, color, estado, timestamp FROM light_sensors WHERE sensor_id = ?",
		sensorID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lights []*entities.LightSensor
	for rows.Next() {
		var light entities.LightSensor
		if err := rows.Scan(&light.ID, &light.SensorID, &light.Intensidad, &light.Color, &light.Estado, &light.Timestamp); err != nil {
			return nil, err
		}
		lights = append(lights, &light)
	}
	return lights, nil
}