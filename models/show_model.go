package models

type Show struct {
	ShowId   uint32 `json:"showId" gorm:"primaryKey"`
	ShowType string `json:"showType"`
	ShowName string `json:"showName"`
}
