package router

import (
	controller "DataConsumer/src/SoundSensor/Infraestructure/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterSoundSensorRoutes(router *gin.Engine, soundSensorController *controller.SoundSensorController) {
	// Creamos un grupo de rutas para el controlador de los sensores de sonido
	soundSensorGroup := router.Group("/soundsensor")
	{
		// Obtener todos los datos de los sensores de sonido
		soundSensorGroup.GET("/", soundSensorController.GetSoundData)

		// Crear un nuevo sensor de sonido
		soundSensorGroup.POST("/", soundSensorController.SaveSoundData)

		// Ruta para WebSocket (para mostrar los datos en tiempo real)
		soundSensorGroup.GET("/ws", soundSensorController.HandleWebSocket)

		
	}
}
