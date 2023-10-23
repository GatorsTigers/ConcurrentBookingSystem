package models

type Movie struct {
	MovieId   uint32 `json:"movieId" gorm:"primaryKey"`
	MovieName string `json:"movieName" gorm:"uniqueIndex;size:100"`
	ImageUrl  string `json:"imageUrl"`
}
