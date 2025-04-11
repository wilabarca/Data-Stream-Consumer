package router

import (
	controller "DataConsumer/src/LightSensor/Infraestructure/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterLightSensorRoutes(router *gin.Engine, lightSensorController *controller.LightController) {
	lightSensorGroup := router.Group("/lightsensor")
	{
		lightSensorGroup.GET("/", lightSensorController.GetLightData)

		lightSensorGroup.POST("/", lightSensorController.SaveLightData)

		lightSensorGroup.GET("/ws/handshake/light", lightSensorController.HandleWebSocket)

		
	}
}