package models

type User struct {
	Email    string `gorm:"primaryKey;not null"`
	UserName string `gorm:"not null"`
	PhoneNo  string `gorm:"not null"`
}
