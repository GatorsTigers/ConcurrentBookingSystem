package models

import "time"

type Ticket struct {
	TicketId uint32 `gorm:"primaryKey"`

	SeatReferId  uint32
	EmailReferId string
	ShowReferId  uint32

	StartTimeRefer         time.Time `gorm:"primaryKey"`
	ScreenCompSchReferName string    `gorm:"primaryKey;index;not null"`
	TheaterCompSchReferId  uint32    `gorm:"primaryKey;index;not null"`

	Seat Seat `json:"-"  gorm:"foreignKey:SeatReferId;References:SeatId"`
	User User `json:"-" gorm:"foreignKey:EmailReferId;References:EmailId"`
	Show Show `json:"-" gorm:"foreignKey:ShowReferId;References:ShowId"`

	ScreenShowSchedule ScreenShowSchedule `json:"-"  gorm:"ForeignKey:StartTimeRefer,ScreenCompSchReferName,TheaterCompSchReferId;References:StartTime,ScreenCompReferName,TheaterCompReferId"`
}
