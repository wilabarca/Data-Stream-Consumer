package router

import (
	controller "DataConsumer/src/SoundSensor/Infraestructure/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterSoundSensorRoutes(router *gin.Engine, soundSensorController *controller.SoundSensorController) {
	soundSensorGroup := router.Group("/soundsensor")
	{
		soundSensorGroup.GET("/", soundSensorController.GetSoundData)

		soundSensorGroup.POST("/", soundSensorController.SaveSoundData)

		soundSensorGroup.GET("/ws/handshake/sound", soundSensorController.HandleWebSocket)

		
	}
}
