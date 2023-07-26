package models

type City struct {
	CityId   int    `json:"cityId" gorm:"primary_key;"`
	CityName string `json:"cityName"`
}
