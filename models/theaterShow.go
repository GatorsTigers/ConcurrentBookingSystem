package models

type TheaterShow struct {
	TheaterReferId int     `json:"theaterId" gorm:"primary_key;"`
	ShowReferId    int     `json:"showId" gorm:"primary_key;"`
	Theater        Theater `gorm:"foreignKey:TheaterReferId;References:TheaterId"`
	Show           Show    `gorm:"foreignKey:ShowReferId;References:ShowId"`
}
