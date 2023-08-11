package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateShows(shows *[]models.Show) error {
	if txn := DbInstance.Db.Create(shows); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetShows(shows *[]models.Show) error {
	if txn := DbInstance.Db.Find(shows); txn.Error != nil {
		return txn.Error
	}
	return nil
}
