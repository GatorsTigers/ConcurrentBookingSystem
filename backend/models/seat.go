package models

type Seat struct {
	SeatId   int    `json:"seatId" gorm:"primary_key;"`
	SeatName string `json:"seatName"`
	ScreenId int    `json:"screenId"`
}
