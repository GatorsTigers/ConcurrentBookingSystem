package models

type ScreenShowSchedule struct {
	StartTime           string `gorm:"primaryKey"`
	ScreenCompReferName string `gorm:"primaryKey;index;not null"`
	TheaterCompReferId  int    `gorm:"primaryKey;index;not null"`
	EndTime             string
	ShowReferId         uint   `gorm:"index;not null"`
	Screen              Screen `gorm:"foreignKey:ScreenCompReferName,TheaterCompReferId;References:ScreenName,TheaterReferId"`
	Show                Show   `gorm:"foreignKey:ShowReferId;References:ShowId"`
}
