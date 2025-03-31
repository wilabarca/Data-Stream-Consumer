package router

import (
	controller "DataConsumer/src/LightSensor/Infraestructure/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterLightSensorRoutes(router *gin.Engine, lightSensorController *controller.LightController) {
	// Creamos un grupo de rutas para el controlador de los sensores de luz
	lightSensorGroup := router.Group("/lightsensor")
	{
		// Obtener todos los sensores de luz
		lightSensorGroup.GET("/", lightSensorController.GetLightData)

		// Crear un nuevo sensor de luz
		lightSensorGroup.POST("/", lightSensorController.SaveLightData)

	   // Ruta para WebSocket (para mostrar los datos en tiempo real)
		lightSensorGroup.GET("/ws", lightSensorController.HandleWebSocket)

		
	}
}