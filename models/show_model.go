package models

type Show struct {
	ShowId   uint32 `gorm:"primaryKey"`
	ShowType string `gorm:"not null"`
	ShowName string `gorm:"not null"`
}
