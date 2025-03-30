package controller

import (
	application "DataConsumer/src/AirQuality/Application"
	entities "DataConsumer/src/AirQuality/Domain/Entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AirQualityController struct {
	service *application.AirQualityService
}

// Constructor del controlador
func NewAirQualityController(service *application.AirQualityService) *AirQualityController {
	return &AirQualityController{
		service: service,
	}
}

// Método para manejar la solicitud HTTP POST para guardar los datos del sensor
func (c *AirQualityController) SaveAirQualityData(ctx *gin.Context) {
	var sensor entities.AirQualitySensor
     
	// Decodificar el cuerpo de la solicitud JSON
	if err := ctx.ShouldBindJSON(&sensor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al decodificar los datos"})
		return
	}

	// Llamar al servicio para guardar los datos
	if err := c.service.SaveAirQualityData(&sensor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos en la base de datos"})
		return
	}

	// Responder con éxito
	ctx.JSON(http.StatusCreated, gin.H{"message": "Datos guardados correctamente"})
}

// Método para manejar la solicitud HTTP GET para obtener todos los datos
func (c *AirQualityController) GetAirQualityData(ctx *gin.Context) {
	// Llamar al servicio para obtener los datos
	data, err := c.service.GetAirQualityData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos"})
		return
	}

	// Responder con los datos en formato JSON
	ctx.JSON(http.StatusOK, data)
}

// Método para manejar las conexiones WebSocket
func (c *AirQualityController) HandleWebSocketConnection(ctx *gin.Context) {
	// (Mantén tu implementación actual de WebSocket aquí)
	// Nota: Para WebSocket necesitarás adaptarlo también al contexto de Gin
}