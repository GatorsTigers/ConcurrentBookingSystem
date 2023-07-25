package models

type seat struct {
	SeatId       int       `json:"seatId" gorm:"primary_key;"`
	SeatName     string    `json:"seatName"`
	ScreenId 	 int       `json:"screenId"`
}