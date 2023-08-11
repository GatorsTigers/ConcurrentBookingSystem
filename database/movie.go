package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateMovies(movies *[]models.Movie) error {
	if txn := DbInstance.Db.Create(movies); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetMovies(movies *[]models.Movie) error {
	if txn := DbInstance.Db.Find(movies); txn.Error != nil {
		return txn.Error
	}
	return nil
}
