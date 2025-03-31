package database

import (
	"database/sql"
	entities "DataConsumer/src/AirQuality/Domain/Entities"
	repositories "DataConsumer/src/AirQuality/Domain/Repositories"
)

type MySQLAirQualityRepository struct {
	db *sql.DB
}

func NewAirQualityRepository(db *sql.DB) repositories.AirQualityRepository {
	return &MySQLAirQualityRepository{db: db}
}

func (m *MySQLAirQualityRepository) SaveAirQualityData(sensor *entities.AirQualitySensor) error {
	_, err := m.db.Exec(
		"INSERT INTO air_quality_sensors (sensor_id, co2_ppm, pm25, temperatura, timestamp) VALUES (?, ?, ?, ?, ?)",
		sensor.SensorID, sensor.CO2PPM, sensor.PM25, sensor.Temperatura, sensor.Timestamp,
	)
	return err
}

func (m *MySQLAirQualityRepository) GetAirQualityData() ([]*entities.AirQualitySensor, error) {
	rows, err := m.db.Query("SELECT id, sensor_id, co2_ppm, pm25, temperatura, timestamp FROM air_quality_sensors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []*entities.AirQualitySensor
	for rows.Next() {
		var sensor entities.AirQualitySensor
		if err := rows.Scan(&sensor.ID, &sensor.SensorID, &sensor.CO2PPM, &sensor.PM25, &sensor.Temperatura, &sensor.Timestamp); err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
	}
	return sensors, nil
}

func (m *MySQLAirQualityRepository) GetAllAirQualityData() ([]*entities.AirQualitySensor, error) {
	return m.GetAirQualityData()
}