package models

import "time"

type ScreenShowSchedule struct {
	StartTime           time.Time `gorm:"primaryKey;not null"`
	EndTime             time.Time `gorm:"not null"`
	ScreenCompReferName string    `gorm:"primaryKey;index;not null"`
	TheaterCompReferId  uint32    `gorm:"primaryKey;index;not null"`
	ShowReferId         uint32    `gorm:"index;not null"`
	Screen              Screen    `gorm:"foreignKey:ScreenCompReferName,TheaterCompReferId;References:ScreenName,TheaterReferId"`
	Show                Show      `gorm:"foreignKey:ShowReferId;References:ShowId"`
}
