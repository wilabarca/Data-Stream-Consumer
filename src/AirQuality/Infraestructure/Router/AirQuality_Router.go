package router

import (
	controller "DataConsumer/src/AirQuality/Infraestructure/Controller"
	"github.com/gin-gonic/gin"
)

// Función para registrar las rutas del sensor de calidad del aire
func RegisterAirQualitySensorRoutes(router *gin.Engine, airQualitySensorController *controller.AirQualityController) {
    // Creamos un grupo de rutas para el controlador
    airQualitySensorGroup := router.Group("/airqualitysensor")
    {
        // Obtener todos los datos
        airQualitySensorGroup.GET("/", airQualitySensorController.GetAirQualityData)

        // Crear nuevos datos
        airQualitySensorGroup.POST("/", airQualitySensorController.SaveAirQualityData)

        // Ruta para WebSocket (para mostrar los datos en tiempo real)
        airQualitySensorGroup.GET("/ws", airQualitySensorController.HandleWebSocketConnection)
    }
}