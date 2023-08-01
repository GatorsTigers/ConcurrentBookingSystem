package models

type ScreenShowSchedule struct {
	StartTime           string `gorm:"primaryKey"`
	ScreenCompReferName string `gorm:"primaryKey;index;not null"`
	TheaterCompReferId  uint32 `gorm:"primaryKey;index;not null"`
	EndTime             string
	ShowReferId         uint32 `gorm:"index;not null"`
	Screen              Screen `json:"-"  gorm:"foreignKey:ScreenCompReferName,TheaterCompReferId;References:ScreenName,TheaterReferId"`
	Show                Show   `json:"-"  gorm:"foreignKey:ShowReferId;References:ShowId"`
}
