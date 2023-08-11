package models

import "time"

type Ticket struct {
	SeatReferId            uint32    `gorm:"primaryKey" json:"seatReferId"`
	EmailReferId           string    `json:"emailReferId"`
	ShowReferId            uint32    `json:"showReferId"`
	StartTimeRefer         time.Time `json:"startTime" gorm:"primaryKey"`
	ScreenCompSchReferName string    `json:"screenName" gorm:"primaryKey;index;not null"`
	TheaterCompSchReferId  uint32    `json:"theaterId" gorm:"primaryKey;index;not null"`

	Seat               Seat               `json:"-"  gorm:"foreignKey:SeatReferId;References:SeatId"`
	User               User               `json:"-" gorm:"foreignKey:EmailReferId;References:EmailId"`
	Show               Show               `json:"-" gorm:"foreignKey:ShowReferId;References:ShowId"`
	ScreenShowSchedule ScreenShowSchedule `json:"-"  gorm:"ForeignKey:StartTimeRefer,ScreenCompSchReferName,TheaterCompSchReferId;References:StartTime,ScreenCompReferName,TheaterCompReferId"`
}
