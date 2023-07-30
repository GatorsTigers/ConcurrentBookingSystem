package main

import (
	"log"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/controller"
	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/gin-gonic/gin"
)

func main() {

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
	router.POST("/addCities", controller.CreateCities)
	router.GET("/showCities", controller.ShowCities)
	router.Run(":8000")
	log.Println("Server running on port 8000")
}
