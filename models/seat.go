package models

type Seat struct {
	SeatId              uint32 `json:"seatId" gorm:"primary_key;"`
	SeatName            string `json:"seatName"`
	ScreenCompReferName string `json:"-" gorm:"not null"`
	TheaterCompReferId  uint32 `json:"-" gorm:"not null"`
	Screen              Screen `json:"-"  gorm:"ForeignKey:ScreenCompReferName,TheaterCompReferId;References:ScreenName,TheaterReferId"`
}
