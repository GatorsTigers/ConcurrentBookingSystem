package models

type Seat struct {
	SeatId              int    `json:"seatId" gorm:"primary_key;"`
	SeatName            string `json:"seatName"`
	ScreenCompReferName string `json:"screenName" gorm:"not null"`
	TheaterCompReferId  int    `json:"theaterId" gorm:"not null"`
	Screen              Screen `gorm:"ForeignKey:ScreenCompReferName,TheaterCompReferId;References:ScreenName,TheaterReferId"`
}
