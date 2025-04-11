package database

import (
	entities "DataConsumer/src/AirQuality/Domain/Entities"
	repositories "DataConsumer/src/AirQuality/Domain/Repositories"
	"database/sql"
	"fmt"
	"time"
)

type MySQLAirQualityRepository struct {
	db *sql.DB
}

func NewAirQualityRepository(db *sql.DB) repositories.AirQualityRepository {
	return &MySQLAirQualityRepository{db: db}
}

func (m *MySQLAirQualityRepository) SaveAirQualityData(sensor *entities.AirQualitySensor) error {
	parsedTime, err := time.Parse(time.RFC3339, sensor.Timestamp)
	if err != nil {
		return fmt.Errorf("error al parsear el timestamp: %v", err)
	}
	mysqlFormattedTime := parsedTime.Format("2006-01-02 15:04:05")

	_, err = m.db.Exec(
		"INSERT INTO airqualitysensor (sensorID, CO2PPM, Air_level, timestamp) VALUES (?, ?, ?, ?)",
		sensor.SensorID, sensor.CO2PPM, sensor.Air_level, mysqlFormattedTime,
	)
	if err != nil {
		return fmt.Errorf("error al insertar datos de calidad del aire: %v", err)
	}
	return nil
}

func (m *MySQLAirQualityRepository) GetAllAirQualityData() ([]*entities.AirQualitySensor, error) {
	rows, err := m.db.Query("SELECT sensorID, CO2PPM, Air_level, timestamp FROM air_quality_sensors")
	if err != nil {
		return nil, fmt.Errorf("error al consultar datos de calidad del aire: %v", err)
	}
	defer rows.Close()

	var sensors []*entities.AirQualitySensor
	for rows.Next() {
		var sensor entities.AirQualitySensor
		if err := rows.Scan(&sensor.SensorID, &sensor.CO2PPM,&sensor.Air_level, &sensor.Timestamp); err != nil {
			return nil, fmt.Errorf("error al escanear datos de calidad del aire: %v", err)
		}
		sensors = append(sensors, &sensor)
	}
	return sensors, nil
}