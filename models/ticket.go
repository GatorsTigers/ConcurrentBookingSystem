package models

// import "time"

type Ticket struct {
	TicketId     uint32 `gorm:"primaryKey" json:"ticketId"`
	SeatReferId  uint32 `json:"seatReferId"`
	EmailReferId string `json:"emailReferId"`
	// MovieReferId           uint32    `json:"movieReferId"`
	ShowReferId uint32 `json:"showReferId"`
	// StartTimeRefer         time.Time `json:"startTime" gorm:"primaryKey"`
	// ScreenCompSchReferName string    `json:"screenName" gorm:"primaryKey;index;not null"`
	// TheaterCompSchReferId  uint32    `json:"theaterId" gorm:"primaryKey;index;not null"`

	Seat Seat `json:"-"  gorm:"foreignKey:SeatReferId;References:SeatId"`
	User User `json:"-" gorm:"foreignKey:EmailReferId;References:EmailId"`
	// Movie Movie `json:"-" gorm:"foreignKey:MovieReferId;References:MovieId"`
	// Show Show `json:"-"  gorm:"ForeignKey:StartTimeRefer,ScreenCompSchReferName,TheaterCompSchReferId;References:StartTime,ScreenCompReferName,TheaterCompReferId"`
	Show Show `json:"-" gorm:"foreignKey:ShowReferId;References:ShowId"`
}
