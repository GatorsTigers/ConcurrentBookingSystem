package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateScreens(screens []models.Screen) ([]models.Screen, error) {
	if txn := DbInstance.Db.Create(&screens); txn.Error != nil {
		return []models.Screen{}, txn.Error
	}
	return screens, nil
}

func ShowScreens() ([]models.Screen, error) {
	var screens []models.Screen
	if txn := DbInstance.Db.Find(&screens); txn.Error != nil {
		return screens, txn.Error
	}
	return screens, nil
}
