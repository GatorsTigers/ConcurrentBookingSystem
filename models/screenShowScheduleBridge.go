package models

type ScreenShowSchedule struct {
	StartTime string `gorm:"primaryKey"`
	EndTime   string
	ScreenId  uint `gorm:"primaryKey;index;not null"`
	ShowId    uint `gorm:"index;not null"`
}
