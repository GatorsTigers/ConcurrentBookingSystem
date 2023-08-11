package models

type TheaterMovie struct {
	TheaterReferId uint32  `json:"theaterId" gorm:"primary_key;"`
	ShowReferId    uint32  `json:"showId" gorm:"primary_key;"`
	Theater        Theater `json:"-"  gorm:"foreignKey:TheaterReferId;References:TheaterId"`
	Movie          Movie   `json:"-"  gorm:"foreignKey:ShowReferId;References:ShowId"`
}
