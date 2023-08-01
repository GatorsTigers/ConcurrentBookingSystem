package controller

import (
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func AddShows(context *gin.Context) {
	var showsRequest []models.Show

	if err := context.BindJSON(&showsRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse city response",
		})
		shows, _ := database.CreateShows(showsRequest)
		context.JSON(http.StatusOK, shows)
	}
}

func GetShows(context *gin.Context) {
	shows, err := database.GetShows()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get cities",
		})
	}
	context.JSON(http.StatusOK, shows)
}
