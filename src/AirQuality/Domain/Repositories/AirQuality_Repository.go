package repositories

import entities "DataConsumer/src/AirQuality/Domain/Entities"

type AirQualityRepository interface {
	SaveAirQualityData(sensor *entities.AirQualitySensor) error
	GetAllAirQualityData() ([]*entities.AirQualitySensor, error)
}