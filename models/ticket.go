package models

import "time"

type Ticket struct {
	TicketId     uint32    `json:"ticketId" gorm:"primary_key;"`
	BookingTime  time.Time `json:"bookingTime"`
	EmailReferId string    `json:"emailReferId"`
	ShowReferId  uint32    `json:"showId"`

	User User `json:"-" gorm:"foreignKey:EmailReferId;References:EmailId"`
	Show Show `json:"-"  gorm:"ForeignKey:ShowReferId;References:ShowId"`
}
