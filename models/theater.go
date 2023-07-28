package models

type Theater struct {
	TheaterId   int    `json:"theaterId" gorm:"primary_key;"`
	TheaterName string `json:"theaterName"`
	CityReferId uint32 `json:"cityReferId"`
	City        City   `gorm:"foreignKey:CityReferId;References:CityId"`
}
