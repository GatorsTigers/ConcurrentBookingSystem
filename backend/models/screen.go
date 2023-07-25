package models

type screen struct {
	ScreenId       int       `json:"screenId" gorm:"primary_key;"`
	ScreenName     string    `json:"screenName"`
	TheaterId 	   int       `json:"theaterId"`
}