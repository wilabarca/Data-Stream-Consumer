package database

import (
	entities "DataConsumer/src/TemperatureHumidity/Domain/Entities"
	repositories "DataConsumer/src/TemperatureHumidity/Domain/Repositories"
	"database/sql"
)

type MySQLTemperatureHumidityRepository struct {
	db *sql.DB
}

// Constructor para el repository
// Constructor para el repository
func NewMySQLTemperatureHumidityRepository(db *sql.DB) repositories.TemperatureHumidityRepository {
    return &MySQLTemperatureHumidityRepository{db: db}
}
// SaveTemperatureHumidityData guarda los datos del sensor de temperatura y humedad en la base de datos.
func (m *MySQLTemperatureHumidityRepository) SaveTemperatureHumidityData(sensor *entities.TemperatureHumiditySensor) error {
	_, err := m.db.Exec("INSERT INTO temperature_humidity_sensors (sensor_id, temperatura, humedad, timestamp) VALUES (?, ?, ?, ?)",
		sensor.SensorID, sensor.Temperatura, sensor.Humedad, sensor.Timestamp)
	return err
}

// GetTemperatureHumidityData obtiene todos los datos de los sensores de temperatura y humedad desde la base de datos.
func (m *MySQLTemperatureHumidityRepository) GetTemperatureHumidityData() ([]*entities.TemperatureHumiditySensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, temperatura, humedad, timestamp FROM temperature_humidity_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var temperatureHumiditySensors []*entities.TemperatureHumiditySensor
	for rows.Next() {
		var sensor entities.TemperatureHumiditySensor
		if err := rows.Scan(&sensor.ID, &sensor.SensorID, &sensor.Temperatura, &sensor.Humedad, &sensor.Timestamp); err != nil {
			return nil, err
		}
		temperatureHumiditySensors = append(temperatureHumiditySensors, &sensor)
	}
	return temperatureHumiditySensors, nil
}
