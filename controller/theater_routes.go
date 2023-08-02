package controller

import (
	"fmt"
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func AddTheaters(context *gin.Context) {
	var theaters []models.Theater
	if err := context.BindJSON(&theaters); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse theater response",
		})
	} else {
		theaterResult, err := database.CreateTheaters(theaters)
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("%s", err),
			})
		} else {
			context.JSON(http.StatusOK, theaterResult)
		}
	}
}

func ShowTheaters(context *gin.Context) {
	theaters, err := database.ShowTheaters()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get cities",
		})
	} else {
		context.JSON(http.StatusOK, theaters)
	}
}