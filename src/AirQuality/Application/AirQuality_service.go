package application

import (
	entities "DataConsumer/src/AirQuality/Domain/Entities"
	repositories "DataConsumer/src/AirQuality/Domain/Repositories"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

// AirQualityService es el servicio que maneja la lógica de negocio relacionada con la calidad del aire.
// Incluye la gestión de clientes WebSocket, almacenamiento de datos y transmisión en tiempo real.
type AirQualityService struct {
	repo      repositories.AirQualityRepository         
	clients   map[*websocket.Conn]bool               
	broadcast chan *entities.AirQualitySensor         
	mu        sync.Mutex                               
}

// NewAirQualityService crea e inicializa un nuevo servicio de calidad del aire.
func NewAirQualityService(repo repositories.AirQualityRepository) *AirQualityService {
	return &AirQualityService{
		repo:      repo,
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.AirQualitySensor),
	}
}

// SaveAirQualityData guarda un dato de calidad del aire y lo transmite a todos los clientes conectados.
func (s *AirQualityService) SaveAirQualityData(sensor *entities.AirQualitySensor) error {
	// Verificar si el dato del sensor es nulo.
	if sensor == nil {
		return errors.New("los datos del sensor de calidad del aire son nulos")
	}

	// Intentar guardar el dato en el repositorio (base de datos o almacenamiento persistente).
	if err := s.repo.SaveAirQualityData(sensor); err != nil {
		return err
	}

	// Enviar los datos al canal broadcast para transmitirlos a través de WebSocket.
	s.broadcast <- sensor
	return nil
}

// GetAirQualityData obtiene todos los datos de calidad del aire desde el repositorio.
func (s *AirQualityService) GetAirQualityData() ([]*entities.AirQualitySensor, error) {
	return s.repo.GetAllAirQualityData()
}

// GetAllAirQualityData también obtiene todos los datos de calidad del aire (duplicado del método anterior).
func (s *AirQualityService) GetAllAirQualityData() ([]*entities.AirQualitySensor, error) {
	return s.repo.GetAllAirQualityData()
}

// HandleWebSocketConnection maneja una nueva conexión WebSocket.
// Permite registrar clientes y gestionar su desconexión.
func (s *AirQualityService) HandleWebSocketConnection(conn *websocket.Conn) {
	// Bloquear la sección crítica para añadir un nuevo cliente.
	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	// Cuando la función termina, se cierra la conexión y se elimina del mapa de clientes.
	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.Close()
	}()

	// Mantiene la conexión abierta mientras el cliente no la cierre o no ocurra un error.
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break // Si ocurre un error al leer el mensaje, se rompe el ciclo y se cierra la conexión.
		}
	}
}

// StartBroadcasting inicia un proceso que transmite datos en tiempo real a todos los clientes conectados.
func (s *AirQualityService) StartBroadcasting() {
	// Este bucle se ejecuta continuamente mientras existan datos enviados al canal broadcast.
	for sensorData := range s.broadcast {
		s.mu.Lock()
		for client := range s.clients {
			// Intenta enviar los datos del sensor a cada cliente WebSocket.
			if err := client.WriteJSON(sensorData); err != nil {
				// Si ocurre un error al enviar datos, se cierra la conexión y se elimina del mapa de clientes.
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()
	}
}
