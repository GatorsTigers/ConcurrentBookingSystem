package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateShows(shows []models.Show) ([]models.Show, error) {
	if txn := DbInstance.Db.Create(&shows); txn.Error != nil {
		return []models.Show{}, txn.Error
	}
	return shows, nil
}

func GetShows() ([]models.Show, error) {
	var shows []models.Show
	if txn := DbInstance.Db.Find(&shows); txn.Error != nil {
		return shows, txn.Error
	}
	return shows, nil
}
