package models

import "time"

type Ticket struct {
	SeatReferId            uint32    `gorm:"primaryKey" json:"seatReferId"`
	EmailReferId           string    `json:"emailReferId"`
	MoviewReferId          uint32    `json:"movieReferId"`
	StartTimeRefer         time.Time `json:"startTime" gorm:"primaryKey"`
	ScreenCompSchReferName string    `json:"screenName" gorm:"primaryKey;index;not null"`
	TheaterCompSchReferId  uint32    `json:"theaterId" gorm:"primaryKey;index;not null"`

	Seat  Seat  `json:"-"  gorm:"foreignKey:SeatReferId;References:SeatId"`
	User  User  `json:"-" gorm:"foreignKey:EmailReferId;References:EmailId"`
	Movie Movie `json:"-" gorm:"foreignKey:MovieReferId;References:MovieId"`
	Show  Show  `json:"-"  gorm:"ForeignKey:StartTimeRefer,ScreenCompSchReferName,TheaterCompSchReferId;References:StartTime,ScreenCompReferName,TheaterCompReferId"`
}
