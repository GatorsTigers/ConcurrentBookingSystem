package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

type ScreenSeats struct {
	ScreenReferId uint32 `json:"screenId"`
	NumSeats      int    `json:"numSeats"`
}

type AddSeatsRequest struct {
	TheaterCompReferId uint32        `json:"theaterId"`
	ScreenSeats        []ScreenSeats `json:"screenSeats"`
}

func AddSeats(context *gin.Context) {
	var request AddSeatsRequest
	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse theater response",
		})
	} else {
		var seats []models.Seat
		for screen := 0; screen < len(request.ScreenSeats); screen++ {
			screenSeat := request.ScreenSeats[screen]

			for seat := 0; seat < screenSeat.NumSeats; seat++ {
				seats = append(seats, models.Seat{
					SeatName:      getSeatName(seat),
					ScreenReferId: screenSeat.ScreenReferId,
				})
			}
			err := database.CreateSeats(&seats)
			if err != nil {
				context.JSON(http.StatusConflict, gin.H{
					"error": fmt.Sprintf("%s", err),
				})
			} else {
				context.JSON(http.StatusOK, seats)
			}
		}

	}
}

func getSeatName(seatNum int) string {
	x := float64(seatNum) / 10
	y := math.Floor(x)
	return string(rune('A')+rune(y)) + strconv.Itoa(seatNum%10)
}
