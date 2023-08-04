package models

type User struct {
	EmailId  string `json:"emailId" gorm:"primaryKey;not null;size:100"`
	Password string `json:"password" gorm:"not null;" validate:"required"`
}
