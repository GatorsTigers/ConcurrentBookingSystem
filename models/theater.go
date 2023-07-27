package models

type Theater struct {
	TheaterId   int    `json:"theaterId" gorm:"primary_key;"`
	TheaterName string `json:"theaterName"`
	CityReferId int    `json:"cityReferId"`
	City        City   `gorm:"foreignKey:CityReferId;References:CityId"`
}
