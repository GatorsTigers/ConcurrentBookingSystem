package models

type Ticket struct {
	TicketId         uint32  `gorm:"primaryKey" json:"ticketId"`
	EmailReferId     string  `json:"emailReferId"`
	Amount           float64 `json:"amount"`
	BankTranactionId string  `json:"transactionId" gorm:"size:50"`
	User             User    `json:"-" gorm:"foreignKey:EmailReferId;References:EmailId"`
}
