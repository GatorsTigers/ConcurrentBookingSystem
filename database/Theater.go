package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateTheaters(theaters []models.Theater) ([]models.Theater, error) {
	db := DbInstance.GetDB()
	if txn := db.Create(&theaters); txn.Error != nil {
		return []models.Theater{}, txn.Error
	}
	return theaters, nil
}

func ShowTheaters() ([]models.Theater, error) {
	var theaters []models.Theater
	db := DbInstance.GetDB()
	if txn := db.Find(&theaters); txn.Error != nil {
		return theaters, txn.Error
	}
	return theaters, nil
}
