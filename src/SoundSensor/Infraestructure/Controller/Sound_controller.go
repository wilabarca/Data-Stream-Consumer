package controller

import (
	application "DataConsumer/src/SoundSensor/Application"
	entities "DataConsumer/src/SoundSensor/Domain/Entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// SoundSensorController es el controlador que maneja las peticiones relacionadas con los sensores de sonido
type SoundSensorController struct {
	service *application.SoundService
	upgrader websocket.Upgrader
}

// NewSoundSensorController es el constructor para el controlador de los sensores de sonido
func NewSoundSensorController(service *application.SoundService) *SoundSensorController {
	return &SoundSensorController{
		service: service,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true }, // Permite cualquier origen
		},
	}
}

// SaveSoundData maneja la creación de nuevos datos del sensor de sonido
func (c *SoundSensorController) SaveSoundData(ctx *gin.Context) {
	var sensor entities.SoundSensor
	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Llamar al servicio para guardar los datos
	err := c.service.SaveSoundData(&sensor)
	if err != nil {
		log.Println("Error saving sound sensor data:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving sound sensor data"})
		return
	}

	// Responder con un mensaje de éxito
	ctx.JSON(http.StatusOK, gin.H{"message": "Sound sensor data saved successfully"})
}

// GetSoundData obtiene todos los datos del sensor de sonido
func (c *SoundSensorController) GetSoundData(ctx *gin.Context) {
	// Llamar al servicio para obtener los datos
	sensors, err := c.service.GetSoundData()
	if err != nil {
		log.Println("Error getting sound sensor data:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting sound sensor data"})
		return
	}

	// Responder con los datos de los sensores
	ctx.JSON(http.StatusOK, sensors)
}

// HandleWebSocket maneja las conexiones WebSocket
func (c *SoundSensorController) HandleWebSocket(ctx *gin.Context) {
	// Establecer la conexión WebSocket
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error upgrading to WebSocket"})
		return
	}
	defer conn.Close()

	// Agregar la conexión al WebSocketHub para manejar múltiples conexiones
	c.service.AddClient(conn)

	// Mantener la conexión abierta y escuchar mensajes del cliente (si es necesario)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading WebSocket message:", err)
			break
		}
	}

	// Eliminar la conexión del WebSocketHub cuando el cliente se desconecta
	c.service.RemoveClient(conn)
}
