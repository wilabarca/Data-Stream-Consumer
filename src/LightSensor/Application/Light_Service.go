package application

import (
	entities "DataConsumer/src/LightSensor/Domain/Entities"
	repositories "DataConsumer/src/LightSensor/Domain/Repositories"
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

type LightService struct {
	repo repositories.LightRepository
	clients map[*websocket.Conn]bool
	broadcast chan *entities.LightSensor
	mu        sync.Mutex
}

// Constructor del servicio para que reciba el repositori como parametro
func NewLightService(repo repositories.LightRepository) *LightService {
	return &LightService{
		repo: repo,
		clients: make(map[*websocket.Conn]bool),
		broadcast: make(chan *entities.LightSensor),
	}
}

// Método para guardar los datos del sensor de luz en la base de datos
// y emitirlos a través del WebSocket si el almacenamiento es exitoso
func (s *LightService) SaveLightData(light *entities.LightSensor) error  {
	// Se verifica si los ddatos son nulos
	if light == nil {
		return errors.New("los datos del sensor de luz son nulos")
	}

	// Guarda los datos a la base de datos a traves del repo
	err := s.repo.SaveLightData(light)
	if err != nil {
		return err
	}

	// se guarda correctamente y enviar los datos al canal de difusion
	s.broadcast <- light

	return nil
}

// Método para obtener todos los datos almacenados del sensor de luz
func (s *LightService) GetLightData() ([]*entities.LightSensor, error)  {
	return s.repo.GetLightData()
}

// Manejo de conexiones WebSocket para enviar datos en tiempo real al frontend
func (s *LightService) HandleWebSocketConnection(conn *websocket.Conn)  {
	// Se bloquea el acesso para modificar el mapa al cliente
	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	// Eliminar el cliente al desconectarse 
	defer func ()  {
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
func (s *LightService) StartBroadcasting() {
	for {
		// Recibir datos del canal de difusión
		lightData := <-s.broadcast

		// Bloquear el acceso al mapa de clientes durante la difusión
		s.mu.Lock()
		for client := range s.clients {
			// Enviar los datos a cada cliente conectado
			err := client.WriteJSON(lightData)
			if err != nil {
				// Si ocurre un error, cerrar la conexión y eliminar el cliente del mapa
				client.Close()
				delete(s.clients, client)
			}
		}
		s.mu.Unlock()
	}
}




