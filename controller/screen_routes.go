package controller

import (
	"fmt"
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func AddScreens(context *gin.Context) {
	var screenJson []models.Screen
	if err := context.BindJSON(&screenJson); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse screen response",
		})
	} else {
		screens, err := database.CreateScreens(screenJson)
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("%s", err),
			})
		} else {
			context.JSON(http.StatusOK, screens)
		}

	}

}

func ShowScreens(context *gin.Context) {
	screens, err := database.ShowScreens()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get screens",
		})
	} else {
		context.JSON(http.StatusOK, screens)
	}

}
