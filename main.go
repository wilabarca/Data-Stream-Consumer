package main

import (
	core "DataConsumer/Core"
	"log"

	airqualityrouter "DataConsumer/src/AirQuality/Infraestructure/Router"
	lightrouter "DataConsumer/src/LightSensor/Infraestructure/Router"
	soundrouter "DataConsumer/src/SoundSensor/Infraestructure/Router"
	temphumidityrouter "DataConsumer/src/TemperatureHumidity/Infraestructure/Router"

	airqualitycontroller "DataConsumer/src/AirQuality/Infraestructure/Controller"
	lightcontroller "DataConsumer/src/LightSensor/Infraestructure/Controller"
	soundcontroller "DataConsumer/src/SoundSensor/Infraestructure/Controller"
	temperaturecontroller "DataConsumer/src/TemperatureHumidity/Infraestructure/Controller"

	airqualityapp "DataConsumer/src/AirQuality/Application"
	lightapp "DataConsumer/src/LightSensor/Application"
	soundapp "DataConsumer/src/SoundSensor/Application"
	temperatureapp "DataConsumer/src/TemperatureHumidity/Application"

	// Importaciones de los paquetes de database
	airqualitydb "DataConsumer/src/AirQuality/Infraestructure/Database"
	lightdb "DataConsumer/src/LightSensor/Infraestructure/Database"
	sounddb "DataConsumer/src/SoundSensor/Infraestructure/Database"
	tempdb "DataConsumer/src/TemperatureHumidity/Infraestructure/Database"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := core.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	// Inicialización de repositorios con las implementaciones concretas
	airQualityRepo := airqualitydb.NewAirQualityRepository(db)
	lightRepo := lightdb.NewLightRepository(db)
	temperatureRepo := tempdb.NewTemperatureHumidityRepository(db)
	soundRepo := sounddb.NewSoundRepository(db)

	// Creación de servicios
	airQualityService := airqualityapp.NewAirQualityService(airQualityRepo)
	lightService := lightapp.NewLightService(lightRepo)
	temperatureService := temperatureapp.NewTemperatureHumidityService(temperatureRepo)
	soundService := soundapp.NewSoundService(soundRepo)

	// Creación de controladores
	airQualityCtrl := airqualitycontroller.NewAirQualityController(airQualityService)
	lightCtrl := lightcontroller.NewLightController(lightService)
	temperatureCtrl := temperaturecontroller.NewTemperatureHumidityController(temperatureService)
	soundCtrl := soundcontroller.NewSoundSensorController(soundService)

	// Registro de rutas
	airqualityrouter.RegisterAirQualitySensorRoutes(router, airQualityCtrl)
	lightrouter.RegisterLightSensorRoutes(router, lightCtrl)
	temphumidityrouter.RegisterTemperatureHumidityRoutes(router, temperatureCtrl)
	soundrouter.RegisterSoundSensorRoutes(router, soundCtrl)

	// Iniciar servidor
	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}