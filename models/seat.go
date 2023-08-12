package models

type Seat struct {
	SeatId            uint32 `json:"seatId" gorm:"primary_key;"`
	SeatName          string `json:"seatName"`
	ScreenCompReferId string `json:"-" gorm:"not null"`
	Screen            Screen `json:"-"  gorm:"ForeignKey:ScreenCompReferId;References:ScreenId"`
}
