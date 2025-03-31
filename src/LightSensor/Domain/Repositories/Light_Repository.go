package repositories

import entities "DataConsumer/src/LightSensor/Domain/Entities"

type LightRepository interface {
    SaveLightData(light *entities.LightSensor) error
    GetLightData() ([]*entities.LightSensor, error)
    GetLightDataBySensorID(sensorID string) ([]*entities.LightSensor, error) // ✔️ Correcta firma
}
