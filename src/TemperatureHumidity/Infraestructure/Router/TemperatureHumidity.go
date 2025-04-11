package router

import (
	controller "DataConsumer/src/TemperatureHumidity/Infraestructure/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterTemperatureHumidityRoutes(router *gin.Engine, temperatureHumidityController *controller.TemperatureHumidityController) {
	temperatureHumidityGroup := router.Group("/temperaturehumidity")
	{
		temperatureHumidityGroup.GET("/", temperatureHumidityController.GetTemperatureHumidityData)

		temperatureHumidityGroup.POST("/", temperatureHumidityController.SaveTemperatureHumidityData)

			temperatureHumidityGroup.GET("/ws/handshake/temperature", temperatureHumidityController.HandleWebSocketConnection)

		
	}
}
