package repositories

import entities "DataConsumer/src/TemperatureHumidity/Domain/Entities"



type TemperatureHumidityRepository interface {
    SaveTemperatureHumidityData(sensor *entities.TemperatureHumiditySensor) error
    GetTemperatureHumidityData() ([]*entities.TemperatureHumiditySensor, error)
}