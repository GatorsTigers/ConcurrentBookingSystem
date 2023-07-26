package models

type TheaterShow struct {
	TheaterId int `json:"theaterId" gorm:"primary_key;"`
	ShowId    int `json:"showId" gorm:"primary_key;"`
}
