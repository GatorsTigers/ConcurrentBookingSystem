package models

type User struct {
	EmailId  string `gorm:"primaryKey;not null;size:100"`
	Password string `json:"password" gorm:"not null;default:null" validate:"required"`
}
