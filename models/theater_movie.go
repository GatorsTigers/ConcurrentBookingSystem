package models

type TheaterMovie struct {
	TheaterReferId uint32  `json:"theaterId" gorm:"primaryKey;"`
	MovieReferId   uint32  `json:"showId" gorm:"primaryKey;"`
	Theater        Theater `json:"-"  gorm:"foreignKey:TheaterReferId;References:TheaterId"`
	Movie          Movie   `json:"-"  gorm:"foreignKey:MovieReferId;References:MovieId"`
}
