package application

import (
	entities "DataConsumer/src/SoundSensor/Domain/Entities"
	repositories "DataConsumer/src/SoundSensor/Domain/Repositories"
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type SoundService struct {
	repo      repositories.SoundSensor
	clients   map[*websocket.Conn]bool
	broadcast chan *entities.SoundSensor
	mu        sync.Mutex
}

func (s *SoundService) RemoveClient(conn *websocket.Conn) {
	panic("unimplemented")
}

func (s *SoundService) AddClient(conn *websocket.Conn) {
	panic("unimplemented")
}

func NewSoundService(repo repositories.SoundSensor) *SoundService {
	return &SoundService{
		repo:      repo,
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.SoundSensor),
	}
}

func (s *SoundService) SaveSoundData(sensor *entities.SoundSensor) error {
	if sensor == nil {
		return errors.New("los datos del sensor de sonido son nulos")
	}

	if err := s.repo.SaveSoundData(sensor); err != nil {
		return err
	}

	s.broadcast <- sensor
	return nil
}

func (s *SoundService) GetSoundData() ([]*entities.SoundSensor, error) {
	return s.repo.GetSoundData()
}

func (s *SoundService) HandleWebSocketConnection(conn *websocket.Conn) {
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

func (s *SoundService) StartBroadcasting() {
	for sensorData := range s.broadcast {
		message, err := json.Marshal(sensorData)
		if err != nil {
			log.Println("Error al serializar datos:", err)
			continue
		}

		s.mu.Lock()
		for client := range s.clients {
			if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("Error al enviar mensaje:", err)
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()
	}
}
