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
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "could not parse theater request",
		})
	} else {
		err := database.CreateTheaters(&theaters)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusConflict, gin.H{
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
	id, _ := strconv.ParseInt(context.Request.URL.Query().Get("cityId"), 10, 32)
	cityId := uint32(id)
	var theaters []models.Theater
	err := database.GetCityTheatres(cityId, &theaters)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get theatres for the city",
		})
	} else {
		context.JSON(http.StatusOK, theaters)
	}
}

/*
func AddMoviesInTheatre(context *gin.Context) {
	var theaterShows []models.TheaterMovie
	if err := context.BindJSON(&theaterShows); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse theater response",
		})
	} else {
		isInserted, err := database.CreateMovieTheaterBridge(&theaterShows)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "could not add shows in theater",
			})
		} else {
			context.JSON(http.StatusOK, isInserted)
		}
	}

}*/

func AddShowsInTheatre(context *gin.Context) {
	var shows []models.Show
	if err := context.BindJSON(&shows); err != nil {
		context.JSON(http.StatusBadRequest, "Couldn't parse screen show schedule request")
	} else {
		err := database.AddShowsInTheatre(&shows)
		//TODO: Add Show_seat
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "could not add shows in theater",
			})
		} else {
			context.JSON(http.StatusOK, shows)
		}
	}

}

func GetShowsForTheatre(context *gin.Context) {
	theaterId, _ := strconv.ParseInt(context.Request.URL.Query().Get("theaterId"), 10, 32)
	theaterReferId := uint32(theaterId)
	var showSeats []models.ShowSeat
	err := database.GetShowsForTheatre(theaterReferId, &showSeats)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get shows in theater",
		})
	} else {
		context.JSON(http.StatusOK, showSeats)
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
		screenSeatMapping[element.ScreenReferId] = append(screenSeatMapping[element.ScreenReferId], models.Seat{
			SeatId:   element.SeatId,
			SeatName: element.SeatName,
		})
	}
	context.JSON(http.StatusOK, screenSeatMapping)
}
