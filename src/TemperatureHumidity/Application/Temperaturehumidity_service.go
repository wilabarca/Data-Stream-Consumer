package application

import (
	entities "DataConsumer/src/TemperatureHumidity/Domain/Entities"
	repositories "DataConsumer/src/TemperatureHumidity/Domain/Repositories"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type TemperatureHumidityService struct {
	repo      repositories.TemperatureHumidityRepository
	clients   map[*websocket.Conn]bool
	broadcast chan *entities.TemperatureHumiditySensor
	mu        sync.Mutex
}

func NewTemperatureHumidityService(repo repositories.TemperatureHumidityRepository) *TemperatureHumidityService {
	return &TemperatureHumidityService{
		repo:      repo,
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.TemperatureHumiditySensor),
	}
}

func (s *TemperatureHumidityService) SaveTemperatureHumidityData(sensor *entities.TemperatureHumiditySensor) error {
	if sensor == nil {
		return errors.New("los datos del sensor de temperatura y humedad son nulos")
	}

	if err := s.repo.SaveTemperatureHumidityData(sensor); err != nil {
		return err
	}

	s.broadcast <- sensor
	return nil
}

func (s *TemperatureHumidityService) GetTemperatureHumidityData() ([]*entities.TemperatureHumiditySensor, error) {
	return s.repo.GetTemperatureHumidityData()
}

func (s *TemperatureHumidityService) HandleWebSocketConnection(conn *websocket.Conn) {
	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.Close()
	}()

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}

func (s *TemperatureHumidityService) StartBroadcasting() {
	for sensorData := range s.broadcast {
		s.mu.Lock()
		for client := range s.clients {
			if err := client.WriteJSON(sensorData); err != nil {
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()
	}
}