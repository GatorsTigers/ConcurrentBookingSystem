package models

type Ticket struct {
	TicketId     uint32 `gorm:"primaryKey"`
	SeatReferId  uint32
	EmailReferId string
	ShowReferId  uint32
	Seat         Seat `json:"-"  gorm:"foreignKey:SeatReferId;References:SeatId"`
	User         User `json:"-" gorm:"foreignKey:EmailReferId;References:EmailId"`
	Show         Show `json:"-" gorm:"foreignKey:ShowReferId;References:ShowId"`

	StartTimeRefer         string `gorm:"primaryKey"`
	ScreenCompSchReferName string `gorm:"primaryKey;index;not null"`
	TheaterCompSchReferId  uint32 `gorm:"primaryKey;index;not null"`

	ScreenShowSchedule ScreenShowSchedule `json:"-"  gorm:"ForeignKey:StartTimeRefer,ScreenCompSchReferName,TheaterCompSchReferId;References:StartTime,ScreenCompReferName,TheaterCompReferId"`
}
