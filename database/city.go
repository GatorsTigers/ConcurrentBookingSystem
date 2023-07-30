package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateCities(cities []models.City) ([]models.City, error) {
	db := DbInstance.GetDB()
	if txn := db.Create(&cities); txn.Error != nil {
		return []models.City{}, txn.Error
	}
	return cities, nil
}

func ShowCities() ([]models.City, error) {
	var cities []models.City
	db := DbInstance.GetDB()
	if txn := db.Find(&cities); txn.Error != nil {
		return cities, txn.Error
	}
	return cities, nil
}
