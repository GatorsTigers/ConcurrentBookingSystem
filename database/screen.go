package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateScreens(screens []models.Screen) ([]models.Screen, error) {
	db := DbInstance.GetDB()
	if txn := db.Create(&screens); txn.Error != nil {
		return []models.Screen{}, txn.Error
	}
	return screens, nil
}

func ShowScreens() ([]models.Screen, error) {
	var screens []models.Screen
	db := DbInstance.GetDB()
	if txn := db.Find(&screens); txn.Error != nil {
		return screens, txn.Error
	}
	return screens, nil
}
