package models

type City struct {
	CityId   uint32 `json:"cityId" gorm:"primaryKey"`
	CityName string `json:"cityName" gorm:"uniqueIndex;size:100"`
	State    string `json:"state"`
	Country  string `json:"country"`
}
