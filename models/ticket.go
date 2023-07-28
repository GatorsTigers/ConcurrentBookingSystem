package models

type Ticket struct {
	TicketId     uint `gorm:"primaryKey"`
	SeatReferId  uint
	EmailReferId string
	ShowReferId  uint
	Seat         Seat `gorm:"foreignKey:SeatReferId;References:SeatId"`
	User         User `gorm:"foreignKey:EmailReferId;References:EmailId"`
	Show         Show `gorm:"foreignKey:ShowReferId;References:ShowId"`

	StartTimeRefer         string `gorm:"primaryKey"`
	ScreenCompSchReferName string `gorm:"primaryKey;index;not null"`
	TheaterCompSchReferId  int    `gorm:"primaryKey;index;not null"`

	ScreenShowSchedule ScreenShowSchedule `gorm:"ForeignKey:StartTimeRefer,ScreenCompSchReferName,TheaterCompSchReferId;References:StartTime,ScreenCompReferName,TheaterCompReferId"`
}
