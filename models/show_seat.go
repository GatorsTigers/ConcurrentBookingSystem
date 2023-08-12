package models

type ShowSeat struct {
	ShowSeatId    uint32 `json:"showSeatId" gorm:"primary_key"`
	Status        string `json:"status"`
	SeatReferId   uint32 `json:"seatId" gorm:"not null"`
	ShowId        uint32 `json:"showId"`
	TicketReferId uint32 `json:"ticketId"`
	Seat          Seat   `json:"-"  gorm:"ForeignKey:SeatReferId;References:SeatId"`
	Show          Show   `json:"-" gorm:"ForeignKey:ShowReferId;References:ShowId"`
	Ticket        Ticket `json:"-" gorm:"ForeignKey:TicketReferId;References:TicketId"`
}
