package repositories

import entities "DataConsumer/src/SoundSensor/Domain/Entities"

type SoundSensor interface {
	SaveSoundData(sensor *entities.SoundSensor) error
	GetSoundData() ([]*entities.SoundSensor, error)
}