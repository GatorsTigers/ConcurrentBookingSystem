package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateShows(shows []models.Show) ([]models.Show, error) {
	db := DbInstance.GetDB()
	if txn := db.Create(&shows); txn.Error != nil {
		return []models.Show{}, txn.Error
	}
	return shows, nil
}

func GetShows() ([]models.Show, error) {
	var shows []models.Show
	db := DbInstance.GetDB()
	if txn := db.Find(&shows); txn.Error != nil {
		return shows, txn.Error
	}
	return shows, nil
}
