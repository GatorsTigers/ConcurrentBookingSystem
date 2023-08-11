package controller

import (
	"fmt"
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func AddShows(context *gin.Context) {
	var showsRequest []models.Show

	if err := context.BindJSON(&showsRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse show response",
		})
	} else {
		err := database.CreateShows(&showsRequest)
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("%s", err),
			})
		} else {
			context.JSON(http.StatusOK, showsRequest)
		}
	}
}

func GetShows(context *gin.Context) {
	var shows []models.Show
	err := database.GetShows(&shows)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get cities",
		})
	} else {
		context.JSON(http.StatusOK, shows)
	}
}
