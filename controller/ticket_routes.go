package controller

import (
	"fmt"
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

func BookTicket(context *gin.Context) {
	var ticket []models.Ticket
	if err := context.BindJSON(&ticket); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("could not parse ticket request %s", err),
		})
	} else {
		err = database.BookSelectedSeats(&ticket)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "could not book ticket",
			})
		} else {
			context.JSON(http.StatusOK, ticket)
		}
	}
}

func GetUserTickets(context *gin.Context) {
	emailId := context.Request.URL.Query().Get("emailId")
	var tickets []models.Ticket
	err := database.GetTicketsForUser(emailId, &tickets)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get tickets",
		})
	} else {
		context.JSON(http.StatusOK, tickets)
	}
}
