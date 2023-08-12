package models

import "time"

type Show struct {
	ShowId            uint32    `json:"showId" gorm:"primaryKey"`
	StartTime         time.Time `json:"startTime" gorm:"not null;type:time"`
	EndTime           time.Time `json:"endTime" gorm:"not null;type:time"`
	ScreenCompReferId string    `json:"screenCompReferId" gorm:"index;not null"`
	MovieReferId      uint32    `json:"movieReferId" gorm:"index;not null"`
	Screen            Screen    `json:"screen" gorm:"foreignKey:ScreenCompReferId,TheaterCompReferId;References:ScreenId,TheaterReferId"`
	Movie             Movie     `json:"movie" gorm:"foreignKey:MovieReferId;References:MovieId"`
}
