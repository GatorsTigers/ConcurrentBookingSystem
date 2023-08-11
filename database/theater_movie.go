package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateMovieTheaterBridge(theaterMovies *[]models.TheaterMovie) (bool, error) {
	if txn := DbInstance.Db.Create(theaterMovies); txn.Error != nil {
		return false, txn.Error
	}
	return true, nil
}
