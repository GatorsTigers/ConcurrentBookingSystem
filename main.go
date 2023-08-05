package main

import (
	"log"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/controller"
	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()
	config := config.GetConfig()
	database.InitDB(config)
	database.DbInstance.CreateTables()
	serveApplication()
}

func serveApplication() {
	router := gin.Default()
	cityGroup := router.Group("/city")
	cityGroup.POST("", controller.ValidateLogin(controller.CreateCities))
	cityGroup.GET("", controller.ValidateLogin(controller.ShowCities))

	theater := router.Group("/theater")
	theater.POST("", controller.AddTheaters)
	theater.GET("", controller.ValidateLogin(controller.ShowTheaters))
	theater.GET("/city", controller.GetTheatresByCity)

	screen := router.Group("/screen")
	screen.POST("", controller.ValidateLogin(controller.AddScreens))
	screen.GET("", controller.ValidateLogin(controller.ShowScreens))

	show := router.Group("/show")
	show.POST("", controller.ValidateLogin(controller.AddShows))
	show.GET("", controller.ValidateLogin(controller.GetShows))

	auth := router.Group("/auth")
	auth.POST("/register", controller.RegisterUser)
	auth.POST("/login", controller.LoginUser)
	auth.POST("/logout", controller.LogoutUser)

	router.Run(":8000")
	log.Println("Server running on port 8000")
}
