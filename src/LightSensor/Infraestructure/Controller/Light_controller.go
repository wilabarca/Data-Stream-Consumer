package controller

import (
	application "DataConsumer/src/LightSensor/Application"
	entities "DataConsumer/src/LightSensor/Domain/Entities"
	repositories "DataConsumer/src/LightSensor/Domain/Repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type LightController struct {
	service *application.LightService
}

func NewLightController(repo repositories.LightRepository) *LightController {
	service := application.NewLightService(repo)
	return &LightController{
		service: service,
	}
}

// Método para manejar la solicitud POST para guardar los datos del sensor de luz
func (c *LightController) SaveLightData(ctx *gin.Context) {
	// Decodificar los datos recibidos en el cuerpo de la solicitud
	var lightData entities.LightSensor
	if err := ctx.ShouldBindJSON(&lightData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer los datos"})
		return
	}

	// Llamar al servicio para guardar los datos
	if err := c.service.SaveLightData(&lightData); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con un mensaje de éxito
	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos de luz guardados exitosamente"})
}

// Método para manejar la solicitud GET para obtener todos los datos del sensor de luz
func (c *LightController) GetLightData(ctx *gin.Context) {
	// Obtener los datos de luz almacenados a través del servicio
	lightData, err := c.service.GetLightData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Responder con los datos de luz en formato JSON
	ctx.JSON(http.StatusOK, lightData)
}

// Método para manejar la solicitud WebSocket para recibir datos en tiempo real
func (c *LightController) HandleWebSocket(ctx *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // Permitir cualquier origen
	}

	// Mejorar la solicitud HTTP a WebSocket
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al establecer la conexión WebSocket"})
		return
	}

	// Llamar al servicio para manejar la conexión WebSocket
	c.service.HandleWebSocketConnection(conn)
}

// Método para iniciar la transmisión de datos a los clientes WebSocket
func (c *LightController) StartBroadcasting() {
	go c.service.StartBroadcasting()
}