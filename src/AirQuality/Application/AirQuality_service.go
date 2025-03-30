package application

import (
	entities "DataConsumer/src/AirQuality/Domain/Entities"
	repositories "DataConsumer/src/AirQuality/Domain/Repositories"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type AirQualityService struct {
	repo repositories.AirQualityRepository
	clients map[*websocket.Conn]bool
	broadcast chan *entities.AirQualitySensor
	mu        sync.Mutex
}


// Constructor del servicio para que reciba el repositorio como parámetro
func NewAirQualityService(repo repositories.AirQualityRepository) *AirQualityService {
	return &AirQualityService{
		repo: repo,
		clients: make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.AirQualitySensor),
	}
}
// Método para guardar los datos del sensor de calidad del aire en la base de datos
// y emitirlos a través del WebSocket si el almacenamiento es exitoso
func (s *AirQualityService) SaveAirQualityData(sensor *entities.AirQualitySensor) error {
	// Se verifica si los datos son nulos
	if sensor == nil {
		return errors.New("los datos del sensor de calidad del aire son nulos")
	}

	// Guarda los datos a la base de datos a través del repositorio
	err := s.repo.SaveAirQualityData(sensor)
	if err != nil {
		return err
	}

	// Los datos se guardaron correctamente y se envían al canal de difusión
	s.broadcast <- sensor

	return nil
}

// Método para obtener todos los datos almacenados del sensor de calidad del aire
func (s *AirQualityService) GetAirQualityData() ([]*entities.AirQualitySensor, error) {
	return s.repo.GetAllAirQualityData()
}

// Manejo de conexiones WebSocket para enviar datos en tiempo real al frontend
func (s *AirQualityService) HandleWebSocketConnection(conn *websocket.Conn) {
	// Se bloquea el acceso para modificar el mapa al cliente
	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	// Eliminar el cliente al desconectarse
	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.Close()
	}()

	// Mantener la conexión abierta mientras el cliente esté conectado
	for {
		// Leer mensajes entrantes (aunque en este caso no se procesan)
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

// Método para transmitir los datos a todos los clientes conectados
func (s *AirQualityService) StartBroadcasting() {
	for {
		// Recibir datos del canal de difusión
		sensorData := <-s.broadcast

		// Bloquear el acceso al mapa de clientes durante la difusión
		s.mu.Lock()
		for client := range s.clients {
			// Enviar los datos a cada cliente conectado
			err := client.WriteJSON(sensorData)
			if err != nil {
				// Si ocurre un error, cerrar la conexión y eliminar el cliente del mapa
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()
	}
}
