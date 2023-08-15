package models

type ShowSeat struct {
	ShowSeatId  uint32 `json:"showSeatId" gorm:"primaryKey"`
	Status      string `json:"status" gorm:"default:'available'"`
	SeatReferId uint32 `json:"seatId" gorm:"not null"`
	ShowReferId uint32 `json:"showId"`
	// TicketReferId uint32 `json:"ticketId"`
	Seat Seat `json:"-"  gorm:"ForeignKey:SeatReferId;References:SeatId"`
	Show Show `json:"-" gorm:"ForeignKey:ShowReferId;References:ShowId"`
	// Ticket        Ticket `json:"-" gorm:"ForeignKey:TicketReferId;References:TicketId"`
}
