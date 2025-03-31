package controller

import (
	application "DataConsumer/src/TemperatureHumidity/Application"
	entities "DataConsumer/src/TemperatureHumidity/Domain/Entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type TemperatureHumidityController struct {
	service  *application.TemperatureHumidityService
	upgrader websocket.Upgrader
}

// NewTemperatureHumidityController crea una nueva instancia del controlador
func NewTemperatureHumidityController(service *application.TemperatureHumidityService) *TemperatureHumidityController {
	return &TemperatureHumidityController{
		service: service,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Permite todas las conexiones (ajusta esto en producción)
				return true
			},
		},
	}
}

// SaveTemperatureHumidityData maneja el POST para guardar datos del sensor
func (c *TemperatureHumidityController) SaveTemperatureHumidityData(ctx *gin.Context) {
	var sensor entities.TemperatureHumiditySensor

	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos: " + err.Error(),
		})
		return
	}

	if err := c.service.SaveTemperatureHumidityData(&sensor); err != nil {
		log.Printf("Error al guardar datos: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudieron guardar los datos",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Datos registrados exitosamente",
		"data":    sensor,
	})
}

// GetTemperatureHumidityData maneja el GET para obtener todos los registros
func (c *TemperatureHumidityController) GetTemperatureHumidityData(ctx *gin.Context) {
	data, err := c.service.GetTemperatureHumidityData()
	if err != nil {
		log.Printf("Error al obtener datos: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudieron recuperar los datos",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count":   len(data),
		"results": data,
	})
}

// HandleWebSocketConnection maneja conexiones WebSocket para actualizaciones en tiempo real
func (c *TemperatureHumidityController) HandleWebSocketConnection(ctx *gin.Context) {
	conn, err := c.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("Error al establecer WebSocket: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "No se pudo iniciar la conexión WebSocket",
		})
		return
	}

	// Registra la conexión en el servicio
	c.service.HandleWebSocketConnection(conn)

	// Nota: La desconexión se maneja automáticamente en el servicio
}