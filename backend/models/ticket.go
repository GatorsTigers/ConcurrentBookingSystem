package models

type Ticket struct {
	TicketId             uint `gorm:"primaryKey"`
	SeatId               uint `gorm:"index;not null"`
	UserId               uint `gorm:"index;not null"`
	ShowId               uint `gorm:"index;not null"`
	ScreenShowScheduleId uint `gorm:"index;not null"`
}
