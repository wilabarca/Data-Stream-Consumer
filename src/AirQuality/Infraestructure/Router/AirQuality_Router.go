package router

import (
	controller "DataConsumer/src/AirQuality/Infraestructure/Controller"
	"github.com/gin-gonic/gin"
)

func RegisterAirQualitySensorRoutes(router *gin.Engine, airQualitySensorController *controller.AirQualityController) {
    airQualitySensorGroup := router.Group("/airqualitysensor")
    {
        airQualitySensorGroup.GET("/", airQualitySensorController.GetAirQualityData)

        airQualitySensorGroup.POST("/", airQualitySensorController.SaveAirQualityData)

        airQualitySensorGroup.GET("/ws/handshake/air", airQualitySensorController.HandleWebSocketConnection)
    }
}