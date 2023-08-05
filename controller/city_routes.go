package controller

import (
	"fmt"
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func CreateCities(context *gin.Context) {
	var cityJson []models.City
	err := context.BindJSON(&cityJson)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("could not parse city request %s", err),
		})
	} else {
		cities, err := database.CreateCities(cityJson)
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"error": "this city already exists",
			})
		} else {
			context.JSON(http.StatusOK, cities)
		}
	}
}

func ShowCities(context *gin.Context) {
	cities, err := database.ShowCities()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get cities",
		})
	} else {
		context.JSON(http.StatusOK, cities)
	}

}
