package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
)

func CreateTheaters(theaters *[]models.Theater) error {
	if txn := DbInstance.Db.Create(theaters); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func ShowTheaters(theaters *[]models.Theater) error {
	if txn := DbInstance.Db.Find(theaters); txn.Error != nil {
		return txn.Error
	}
	return nil
}

func GetCityTheatres(cityName string, theaters *[]models.Theater) error {
	if txn := DbInstance.Db.Where("city_refer_name = ?", cityName).Find(theaters); txn.Error != nil {
		return txn.Error
	}
	return nil
}
