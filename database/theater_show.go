package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateShowTheaterBridge(theaterShows *[]models.TheaterShow) (bool, error) {
	if txn := DbInstance.Db.Create(theaterShows); txn.Error != nil {
		return false, txn.Error
	}
	return true, nil
}
