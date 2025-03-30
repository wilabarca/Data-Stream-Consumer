package application

import (
	"encoding/json"
	"DataConsumer/src/SoundSensor/Domain/Entities"
	"DataConsumer/src/SoundSensor/Domain/Repositories"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

// WebSocketHub maneja todas las conexiones WebSocket activas.
type WebSocketHub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

// SoundSensorService es la estructura que contiene el repositorio de sensores de sonido
type SoundSensorService struct {
	repository   repositories.SoundSensor
	websocketHub *WebSocketHub
}

// NewSoundSensorService es el constructor para el servicio de sensores de sonido
func NewSoundSensorService(repository repositories.SoundSensor, websocketHub *WebSocketHub) *SoundSensorService {
	return &SoundSensorService{
		repository:   repository,
		websocketHub: websocketHub,
	}
}

// SaveSoundData guarda los datos de un sensor de sonido en la base de datos y envía los datos por WebSocket
func (s *SoundSensorService) SaveSoundData(sensor *entities.SoundSensor) error {
	// Guardar los datos del sensor en la base de datos
	err := s.repository.SaveSoundData(sensor)
	if err != nil {
		return err
	}

	// Enviar los datos al frontend a través de WebSocket
	err = s.sendDataToWebSocket(sensor)
	if err != nil {
		log.Println("Error al enviar los datos por WebSocket:", err)
		return err
	}

	return nil
}

// GetAllSoundData obtiene todos los datos de los sensores de sonido desde la base de datos
func (s *SoundSensorService) GetAllSoundData() ([]*entities.SoundSensor, error) {
	// Llamar al repositorio para obtener todos los datos
	soundSensors, err := s.repository.GetSoundData()
	if err != nil {
		return nil, err
	}

	// Enviar los datos al frontend a través de WebSocket
	for _, sensor := range soundSensors {
		err := s.sendDataToWebSocket(sensor)
		if err != nil {
			log.Println("Error al enviar los datos por WebSocket:", err)
			return nil, err
		}
	}

	return soundSensors, nil
}

// sendDataToWebSocket envía los datos del sensor a través de WebSocket al cliente
func (s *SoundSensorService) sendDataToWebSocket(sensor *entities.SoundSensor) error {
	// Convertir el objeto del sensor a JSON
	message, err := json.Marshal(sensor)
	if err != nil {
		return err
	}

	// Enviar el mensaje a todos los clientes conectados
	s.websocketHub.mu.Lock()
	defer s.websocketHub.mu.Unlock()
	for client := range s.websocketHub.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error al enviar mensaje WebSocket:", err)
			client.Close()
			delete(s.websocketHub.clients, client)
		}
	}

	return nil
}

// AddClient agrega un cliente al WebSocketHub
func (s *SoundSensorService) AddClient(client *websocket.Conn) {
	s.websocketHub.mu.Lock()
	defer s.websocketHub.mu.Unlock()
	s.websocketHub.clients[client] = true
}

// RemoveClient elimina un cliente del WebSocketHub
func (s *SoundSensorService) RemoveClient(client *websocket.Conn) {
	s.websocketHub.mu.Lock()
	defer s.websocketHub.mu.Unlock()
	delete(s.websocketHub.clients, client)
}
