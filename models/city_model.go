package models

type City struct {
	CityId   uint32 `json:"cityId" gorm:"primary_key;"`
	CityName string `json:"cityName"`
}
