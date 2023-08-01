package models

type TheaterShow struct {
	TheaterReferId uint32  `json:"theaterId" gorm:"primary_key;"`
	ShowReferId    uint32  `json:"showId" gorm:"primary_key;"`
	Theater        Theater `json:"-"  gorm:"foreignKey:TheaterReferId;References:TheaterId"`
	Show           Show    `json:"-"  gorm:"foreignKey:ShowReferId;References:ShowId"`
}
