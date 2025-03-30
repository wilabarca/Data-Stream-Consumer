package database

import (
	entities "DataConsumer/src/AirQuality/Domain/Entities"
	repositories "DataConsumer/src/AirQuality/Domain/Repositories"
	"database/sql"
)

type MySQLAirQualityRepository struct {
	db *sql.DB
}

// GetAllAirQualityData implements repositories.AirQualityRepository.
func (m *MySQLAirQualityRepository) GetAllAirQualityData() ([]*entities.AirQualitySensor, error) {
	panic("unimplemented")
}

// Constructor para el repository
func NeeMySQLAirQualityRepository(db *sql.DB) repositories.AirQualityRepository {
	return &MySQLAirQualityRepository{db: db}
}

// SaveAirQualityData guarda los datos del sensor de calidad del aire en la base de datos.
func (m *MySQLAirQualityRepository) SaveAirQualityData(sensor *entities.AirQualitySensor) error {
	_, err := m.db.Exec("INSERT INTO air_quality_sensors (sensor_id, co2_ppm, pm25, temperatura, timestamp) VALUES (?, ?, ?, ?, ?)",
		sensor.SensorID, sensor.CO2PPM, sensor.PM25, sensor.Temperatura, sensor.Timestamp)
	return err
}

// GetAirQualityData obtiene todos los datos de los sensores de calidad del aire desde la base de datos.
func (m *MySQLAirQualityRepository) GetAirQualityData() ([]*entities.AirQualitySensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, co2_ppm, pm25, temperatura, timestamp FROM air_quality_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var airQualitySensors []*entities.AirQualitySensor
	for rows.Next() {
		var sensor entities.AirQualitySensor
		if err := rows.Scan(&sensor.ID, &sensor.SensorID, &sensor.CO2PPM, &sensor.PM25, &sensor.Temperatura, &sensor.Timestamp); err != nil {
			return nil, err
		}
		airQualitySensors = append(airQualitySensors, &sensor)
	}
	return airQualitySensors, nil
}
