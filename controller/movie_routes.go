package controller

import (
	"fmt"
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func AddMovies(context *gin.Context) {
	var moviesRequest []models.Movie

	if err := context.BindJSON(&moviesRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse show response",
		})
	} else {
		err := database.CreateMovies(&moviesRequest)
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("%s", err),
			})
		} else {
			context.JSON(http.StatusOK, moviesRequest)
		}
	}
}

func GetMovies(context *gin.Context) {
	var movies []models.Movie
	err := database.GetMovies(&movies)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get cities",
		})
	} else {
		context.JSON(http.StatusOK, movies)
	}
}
