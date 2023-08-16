package controller

import (
	"fmt"
	"net/http"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookSeatRequest struct {
	EmailId     string // consider using it from cookie/session
	ShowSeatIds []uint32
}

func BookTicket(context *gin.Context) {
	var request BookSeatRequest
	ticket := &models.Ticket{
		EmailReferId:     request.EmailId,
		Amount:           database.GetTotalBookingAmount(request.ShowSeatIds),
		BankTranactionId: uuid.NewString(),
	}
	if err := context.BindJSON(&request); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("could not parse booking request %s", err),
		})
	} else {
		err = database.BookSelectedSeats(ticket)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "could not book ticket",
			})
		} else {
			database.UpdateShowSeats(ticket.TicketId, request.ShowSeatIds)
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
