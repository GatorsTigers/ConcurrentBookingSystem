package main

import (
	"log"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/controller"
	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Successfullty Inserted")
	config := config.GetConfig()
	database.InitDB(config)
	client := database.DbInstance.GetDB()
	database.CreateTables(client)
	// database.InsertDataIntoTables(client)
	log.Printf("Successfullty Inserted")
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	cityGroup := router.Group("/city")
	cityGroup.POST("", controller.CreateCities)
	cityGroup.GET("", controller.ShowCities)

	theater := router.Group("/theater")
	theater.POST("", controller.AddTheaters)
	theater.GET("", controller.ShowTheaters)

	screen := router.Group("/screen")
	screen.POST("", controller.AddScreens)
	screen.GET("", controller.ShowScreens)

	show := router.Group("/show")
	show.POST("", controller.AddShows)
	show.GET("", controller.GetShows)

	router.Run(":8000")
	log.Println("Server running on port 8000")
}
