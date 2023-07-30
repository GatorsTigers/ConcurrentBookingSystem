package models

type City struct {
	CityName string `json:"cityName" gorm:"primary_key;size:100"`
	State    string `json:"state"`
	Country  string `json:"country"`
}
