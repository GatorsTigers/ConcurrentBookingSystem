package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateTheaters(theaters []models.Theater) ([]models.Theater, error) {
	if txn := DbInstance.Db.Create(&theaters); txn.Error != nil {
		return []models.Theater{}, txn.Error
	}
	return theaters, nil
}

func ShowTheaters() ([]models.Theater, error) {
	var theaters []models.Theater
	if txn := DbInstance.Db.Find(&theaters); txn.Error != nil {
		return theaters, txn.Error
	}
	return theaters, nil
}

func GetCityTheatres(cityName string) ([]models.Theater, error) {
	var theaters []models.Theater
	if txn := DbInstance.Db.Where("city_refer_name = ?", cityName).Find(&theaters); txn.Error != nil {
		return theaters, txn.Error
	}
	return theaters, nil
}
