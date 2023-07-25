package model

type Show struct {
	ShowId   uint   `gorm:"primaryKey"`
	ShowType string `gorm:"not null"`
	ShowName string `gorm:"not null"`
}