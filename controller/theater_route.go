package controller

import (
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func AddTheaters(context *gin.Context) {
	var theaters []models.Theater
	if err := context.BindJSON(&theaters); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse city response",
		})
	}
	cities, _ := database.CreateTheaters(theaters)
	context.JSON(http.StatusOK, cities)
}

func ShowTheaters(context *gin.Context) {
	cities, err := database.ShowTheaters()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get cities",
		})
	}
	context.JSON(http.StatusOK, cities)
}
