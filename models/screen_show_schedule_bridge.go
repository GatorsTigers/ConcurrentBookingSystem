package models

import "time"

type ScreenShowSchedule struct {
	StartTime           time.Time `json:"startTime" gorm:"primaryKey;not null;type:time"`
	EndTime             time.Time `json:"endTime" gorm:"not null;type:time"`
	ScreenCompReferName string    `json:"screenCompReferName" gorm:"primaryKey;index;not null"`
	TheaterCompReferId  uint32    `json:"theaterCompReferId" gorm:"primaryKey;index;not null"`
	MovieReferId        uint32    `json:"movieReferId" gorm:"index;not null"`
	Screen              Screen    `json:"screen" gorm:"foreignKey:ScreenCompReferName,TheaterCompReferId;References:ScreenName,TheaterReferId"`
	Movie               Movie     `json:"movie" gorm:"foreignKey:MovieReferId;References:MovieId"`
}
