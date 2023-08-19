package models

type Theater struct {
	TheaterId   uint32 `json:"theaterId" gorm:"primaryKey;"`
	TheaterName string `json:"theaterName" gorm:"uniqueIndex;size:100;"`
	CityReferId uint32 `json:"cityReferId"`
	City        City   `json:"-" gorm:"foreignKey:CityReferId;ReferencesName:CityId"`
}
