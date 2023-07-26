package models

type Theater struct {
	TheaterId   int    `json:"theaterId" gorm:"primary_key;"`
	TheaterName string `json:"theaterName"`
	CityId      int    `json:"cityId"`
}
