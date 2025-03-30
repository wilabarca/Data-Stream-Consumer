package controller

import (
	application "DataConsumer/src/TemperatureHumidity/Application"
	entities "DataConsumer/src/TemperatureHumidity/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemperatureHumidityController struct {
	service *application.TemperatureHumidityService
}

// Constructor del controlador
func NewTemperatureHumidityController(service *application.TemperatureHumidityService) *TemperatureHumidityController {
	return &TemperatureHumidityController{
		service: service,
	}
}
// Método para manejar la solicitud HTTP POST para guardar los datos del sensor
func (c *TemperatureHumidityController) SaveTemperatureHumidityData(ctx *gin.Context) {
	var sensor entities.TemperatureHumiditySensor

	// Decodificar el cuerpo de la solicitud JSON
	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos"})
		return
	}

	// Llamar al servicio para guardar los datos
	if err := c.service.SaveTemperatureHumidityData(&sensor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos en la base de datos"})
		return
	}

	// Responder con éxito
	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente"})
}

// Método para manejar la solicitud HTTP GET para obtener todos los datos
func (c *TemperatureHumidityController) GetTemperatureHumidityData(ctx *gin.Context) {
	// Llamar al servicio para obtener los datos
	data, err := c.service.GetTemperatureHumidityData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos"})
		return
	}

	// Responder con los datos en formato JSON
	ctx.JSON(http.StatusOK, data)
}

// Método para manejar las conexiones WebSocket
func (c *TemperatureHumidityController) HandleWebSocketConnection(ctx *gin.Context) {
	// Aquí puedes mantener la implementación de WebSocket, adaptada para este contexto
	// Nota: Para WebSocket necesitarás adaptarlo también al contexto de Gin
}