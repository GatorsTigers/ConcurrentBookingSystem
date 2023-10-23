package main

import (
	"log"
	"time"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/controller"
	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/logger"
	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With, Content-Type, Accept, Authorization"},
		ExposeHeaders:    []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	cityGroup := router.Group("/city")
	cityGroup.POST("", controller.CreateCities)
	cityGroup.GET("", controller.ShowCities)

	theater := router.Group("/theater")
	theater.POST("", controller.AddTheaters)
	theater.GET("", controller.ShowTheaters)
	theater.GET("/city", controller.GetTheatresByCity)
	theater.POST("/addShow", controller.AddShowsInTheatre)
	theater.GET("/getShow", controller.GetShowsForTheatre)
	theater.GET("/getSeats", controller.GetSeatsForTheater)

	screen := router.Group("/screen")
	screen.POST("", controller.AddScreens)
	screen.GET("", controller.ShowScreens)

	movie := router.Group("/movie")
	movie.POST("", controller.AddMovies)
	movie.GET("", controller.GetMovies)

	seats := router.Group("/seat")
	seats.POST("", controller.AddSeats)

	auth := router.Group("/auth")
	auth.POST("/register", controller.RegisterUser)
	auth.POST("/login", controller.LoginUser)
	auth.POST("/logout", controller.LogoutUser)

	ticket := router.Group("/ticket")
	ticket.POST("", controller.BookTicket)
	ticket.GET("", controller.GetUserTickets)

	router.Run(":8000")
	log.Println("Server running on port 8000")
}
