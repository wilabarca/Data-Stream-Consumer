package router

import (
	controller "DataConsumer/src/TemperatureHumidity/Infraestructure/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterTemperatureHumidityRoutes(router *gin.Engine, temperatureHumidityController *controller.TemperatureHumidityController) {
	// Creamos un grupo de rutas para el controlador de los sensores de temperatura y humedad
	temperatureHumidityGroup := router.Group("/temperaturehumidity")
	{
		// Obtener todos los datos de los sensores de temperatura y humedad
		temperatureHumidityGroup.GET("/", temperatureHumidityController.GetTemperatureHumidityData)

		// Crear un nuevo sensor de temperatura y humedad
		temperatureHumidityGroup.POST("/", temperatureHumidityController.SaveTemperatureHumidityData)

		// Ruta para WebSocket (para mostrar los datos en tiempo real)
		temperatureHumidityGroup.GET("/ws", temperatureHumidityController.HandleWebSocketConnection)

		
	}
}
