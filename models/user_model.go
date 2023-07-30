package models

type User struct {
	EmailId  string `gorm:"primaryKey;not null;size:100"`
	PhoneNo  string `gorm:"not null"`
	Password string
}
