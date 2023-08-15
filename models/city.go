package models

type City struct {
	CityId   uint32 `json:"cityId" gorm:"primaryKey"`
	CityName string `json:"cityName" gorm:"index:city_state,unique;size:100"`
	State    string `json:"state" gorm:"index:city_state,unique;size:100"`
	Country  string `json:"country"`
}
