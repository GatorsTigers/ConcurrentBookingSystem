package controller

import (
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func CreateCities(context *gin.Context) {
	var cityJson []models.City
	if err := context.BindJSON(&cityJson); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse city response",
		})
	}
	cities, _ := database.CreateCities(cityJson)
	context.JSON(http.StatusOK, cities)
}

func ShowCities(context *gin.Context) {
	cities, err := database.ShowCities()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get cities",
		})
	}
	context.JSON(http.StatusOK, cities)
}
