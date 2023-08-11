package models

import "time"

type ScreenShowSchedule struct {
	StartTime           time.Time `json:"startTime" gorm:"primaryKey;not null;type:time"`
	EndTime             time.Time `json:"endTime" gorm:"not null;type:time"`
	ScreenCompReferName string    `json:"screenCompReferName" gorm:"primaryKey;index;not null"`
	TheaterCompReferId  uint32    `json:"theaterCompReferId" gorm:"primaryKey;index;not null"`
	ShowReferId         uint32    `json:"showReferId" gorm:"index;not null"`
	Screen              Screen    `json:"screen" gorm:"foreignKey:ScreenCompReferName,TheaterCompReferId;References:ScreenName,TheaterReferId"`
	Show                Show      `json:"show" gorm:"foreignKey:ShowReferId;References:ShowId"`
}
