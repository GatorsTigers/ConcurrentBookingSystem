package controller

import (
	"fmt"
	"net/http"
	"strconv"

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
		err := database.CreateTheaters(&theaters)
		if err != nil {
			context.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("%s", err),
			})
		} else {
			context.JSON(http.StatusOK, theaters)
		}
	}
}

func ShowTheaters(context *gin.Context) {
	var theaters []models.Theater
	err := database.ShowTheaters(&theaters)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get theatres",
		})
	} else {
		context.JSON(http.StatusOK, theaters)
	}
}

func GetTheatresByCity(context *gin.Context) {
	cityName := context.Request.URL.Query().Get("cityName")
	var theaters []models.Theater
	err := database.GetCityTheatres(cityName, &theaters)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get theatres for the city",
		})
	} else {
		context.JSON(http.StatusOK, theaters)
	}
}

func AddShowsInTheatre(context *gin.Context) {
	var theaterShows []models.TheaterShow
	if err := context.BindJSON(&theaterShows); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse theater response",
		})
	} else {
		isInserted, err := database.CreateShowTheaterBridge(&theaterShows)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "could not add shows in theater",
			})
		} else {
			context.JSON(http.StatusOK, isInserted)
		}
	}

}

func AddScreenShowScheduleInTheatre(context *gin.Context) {
	var screenShowSchedules []models.ScreenShowSchedule
	if err := context.BindJSON(&screenShowSchedules); err != nil {
		context.JSON(http.StatusBadRequest, "Couldn't parse screen show schedule request")
	} else {
		err := database.AddScreenShowScheduleInTheatre(&screenShowSchedules)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "could not add shows in theater",
			})
		} else {
			context.JSON(http.StatusOK, screenShowSchedules)
		}
	}

}

func GetShowsForTheatre(context *gin.Context) {
	theaterId, _ := strconv.ParseInt(context.Request.URL.Query().Get("theaterId"), 10, 32)
	theaterReferId := uint32(theaterId)
	var screenShowSchedules []models.ScreenShowSchedule
	err := database.GetShowScheduleForTheatre(theaterReferId, &screenShowSchedules)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get shows in theater",
		})
	} else {
		context.JSON(http.StatusOK, screenShowSchedules)
	}
}

func GetSeatsForTheater(context *gin.Context) {
	theaterId, _ := strconv.ParseInt(context.Request.URL.Query().Get("theaterId"), 10, 32)
	theaterReferId := uint32(theaterId)
	var screenSeats []models.Seat
	if err := database.GetSeats(int(theaterReferId), &screenSeats); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, fmt.Sprintf("could not get seats for theater %s", err))
	}

	screenSeatMapping := make(map[string][]models.Seat)
	for _, element := range screenSeats {
		screenSeatMapping[element.ScreenCompReferName] = append(screenSeatMapping[element.ScreenCompReferName], models.Seat{
			SeatId:   element.SeatId,
			SeatName: element.SeatName,
		})
	}
	context.JSON(http.StatusOK, screenSeatMapping)
}
