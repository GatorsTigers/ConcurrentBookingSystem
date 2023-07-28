package models

type Theater struct {
	TheaterId   uint32 `json:"theaterId" gorm:"primary_key;"`
	TheaterName string `json:"theaterName"`
	CityReferId uint32 `json:"cityReferId"`
	City        City   `gorm:"foreignKey:CityReferId;References:CityId"`
}
