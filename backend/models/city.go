package models

type theater struct {
	CityId       int       `json:"cityId" gorm:"primary_key;"`
	CityName     string     `json:"cityName"`
}