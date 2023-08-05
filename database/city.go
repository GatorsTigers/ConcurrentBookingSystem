package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateCities(cities []models.City) ([]models.City, error) {
	if txn := DbInstance.Db.Create(&cities); txn.Error != nil {
		return nil, txn.Error
	}
	return cities, nil
}

func ShowCities() ([]*models.City, error) {
	var cities []*models.City
	if txn := DbInstance.Db.Find(&cities); txn.Error != nil {
		return nil, txn.Error
	}
	return cities, nil
}
