package application

import (
	entities "DataConsumer/src/TemperatureHumidity/Domain/Entities"
	repositories "DataConsumer/src/TemperatureHumidity/Domain/Repositories"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type TemperatureHumidityService struct {
	repo repositories.TemperatureHumidityRepository
	clients map[*websocket.Conn]bool
	broadcast chan *entities.TemperatureHumiditySensor
	mu        sync.Mutex
}

// Constructor del servicio para que reciba el repositorio como parámetro
func NewTemperatureHumidityService(repo repositories.TemperatureHumidityRepository) *TemperatureHumidityService {
	return &TemperatureHumidityService{
		repo:      repo,
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.TemperatureHumiditySensor),
	}
}

// Método para guardar los datos del sensor de temperatura y humedad en la base de datos
// y emitirlos a través del WebSocket si el almacenamiento es exitoso
func (s *TemperatureHumidityService) SaveTemperatureHumidityData(sensor *entities.TemperatureHumiditySensor) error {
	// Se verifica si los datos son nulos
	if sensor == nil {
		return errors.New("los datos del sensor de temperatura y humedad son nulos")
	}

	// Guarda los datos a la base de datos a través del repositorio
	err := s.repo.SaveTemperatureHumidityData(sensor)
	if err != nil {
		return err
	}

	// Los datos se guardaron correctamente y se envían al canal de difusión
	s.broadcast <- sensor

	return nil
}

// Método para obtener todos los datos almacenados del sensor de temperatura y humedad
func (s *TemperatureHumidityService) GetTemperatureHumidityData() ([]*entities.TemperatureHumiditySensor, error) {
	return s.repo.GetTemperatureHumidityData()
}

// Manejo de conexiones WebSocket para enviar datos en tiempo real al frontend
func (s *TemperatureHumidityService) HandleWebSocketConnection(conn *websocket.Conn) {
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
func (s *TemperatureHumidityService) StartBroadcasting() {
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