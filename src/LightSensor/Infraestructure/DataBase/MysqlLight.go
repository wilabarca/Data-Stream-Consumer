package database

import (
	entities "DataConsumer/src/LightSensor/Domain/Entities"
	repositories "DataConsumer/src/LightSensor/Domain/Repositories"
	"database/sql"
)

type MySQLSensorRepository struct {
	db *sql.DB
}

func NewMySQLSensorRepository(db *sql.DB) repositories.LightRepository {
	return &MySQLSensorRepository{db: db}
	
}

// SaveLightData guarda los datos del sensor de luz en la base de datos.
func (m *MySQLSensorRepository) SaveLightData(light *entities.LightSensor) error {
	_, err := m.db.Exec("INSERT INTO light_sensors (sensor_id, intensidad, color, estado, timestamp) VALUES (?, ?, ?, ?, ?)", 
		light.SensorID, light.Intensidad, light.Color, light.Estado, light.Timestamp)
	return err
}
// GetLightData obtiene todos los datos de los sensores de luz desde la base de datos.
func (m *MySQLSensorRepository) GetLightData() ([]*entities.LightSensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, intensidad, color, estado, timestamp FROM light_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lightSensors []*entities.LightSensor
	for rows.Next() {
		var light entities.LightSensor
		if err := rows.Scan(&light.ID, &light.SensorID, &light.Intensidad, &light.Color, &light.Estado, &light.Timestamp); err != nil {
			return nil, err
		}
		lightSensors = append(lightSensors, &light)
	}
	return lightSensors, nil
}

// GetLightDataBySensorID obtiene los datos de los sensores de luz por su sensor_id.
func (m *MySQLSensorRepository) GetLightDataBySensorID(sensorID string) ([]*entities.LightSensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, intensidad, color, estado, timestamp FROM light_sensors WHERE sensor_id = ?", sensorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lightSensors []*entities.LightSensor
	for rows.Next() {
		var light entities.LightSensor
		if err := rows.Scan(&light.ID, &light.SensorID, &light.Intensidad, &light.Color, &light.Estado, &light.Timestamp); err != nil {
			return nil, err
		}
		lightSensors = append(lightSensors, &light)
	}
	return lightSensors, nil
}