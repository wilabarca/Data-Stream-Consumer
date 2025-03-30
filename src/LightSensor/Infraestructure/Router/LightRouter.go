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

		// Ruta para WebSocket
		lightSensorGroup.GET("/ws", lightSensorController.HandleWebSocket)

		// No es necesario actualizar los datos de un sensor específico, pero si lo necesitaras:
		// lightSensorGroup.PUT("/:sensor_id", lightSensorController.UpdateLightSensorData)

		// Eliminar un sensor de luz por su ID
		// lightSensorGroup.DELETE("/:sensor_id", lightSensorController.DeleteLightSensorData)
	}
}