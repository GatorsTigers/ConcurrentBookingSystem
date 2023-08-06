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

func GetShows() ([]models.Show, error) {
	var shows []models.Show
	if txn := DbInstance.Db.Find(&shows); txn.Error != nil {
		return shows, txn.Error
	}
	return shows, nil
}
