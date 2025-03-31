package application

import (
	entities "DataConsumer/src/LightSensor/Domain/Entities"
	repositories "DataConsumer/src/LightSensor/Domain/Repositories"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type LightService struct {
	repo      repositories.LightRepository
	clients   map[*websocket.Conn]bool
	broadcast chan *entities.LightSensor
	mu        sync.Mutex
}

func NewLightService(repo repositories.LightRepository) *LightService {
	return &LightService{
		repo:      repo,
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.LightSensor),
	}
}

func (s *LightService) SaveLightData(light *entities.LightSensor) error {
	if light == nil {
		return errors.New("los datos del sensor de luz son nulos")
	}

	if err := s.repo.SaveLightData(light); err != nil {
		return err
	}

	s.broadcast <- light
	return nil
}

func (s *LightService) GetLightData() ([]*entities.LightSensor, error) {
	return s.repo.GetLightData()
}

// 🔨 Corregido para devolver una lista de LightSensor
func (s *LightService) GetLightDataBySensorID(sensorID string) ([]*entities.LightSensor, error) {
	data, err := s.repo.GetLightDataBySensorID(sensorID)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("no data found for the given sensor ID")
	}
	return data, nil
}

func (s *LightService) HandleWebSocketConnection(conn *websocket.Conn) {
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

func (s *LightService) StartBroadcasting() {
	for lightData := range s.broadcast {
		s.mu.Lock()
		for client := range s.clients {
			if err := client.WriteJSON(lightData); err != nil {
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()
	}
}
