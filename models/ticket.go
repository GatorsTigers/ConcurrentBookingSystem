package models

type Ticket struct {
	TicketId     uint32 `gorm:"primaryKey"`
	SeatReferId  uint32
	EmailReferId string
	ShowReferId  uint32
	Seat         Seat `gorm:"foreignKey:SeatReferId;References:SeatId"`
	User         User `gorm:"foreignKey:EmailReferId;References:EmailId"`
	Show         Show `gorm:"foreignKey:ShowReferId;References:ShowId"`

	StartTimeRefer         string `gorm:"primaryKey"`
	ScreenCompSchReferName string `gorm:"primaryKey;index;not null"`
	TheaterCompSchReferId  uint32 `gorm:"primaryKey;index;not null"`

	ScreenShowSchedule ScreenShowSchedule `gorm:"ForeignKey:StartTimeRefer,ScreenCompSchReferName,TheaterCompSchReferId;References:StartTime,ScreenCompReferName,TheaterCompReferId"`
}
