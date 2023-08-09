package models

type Screen struct {
	ScreenName     string  `json:"screenName" gorm:"primary_key;size:100"`
	TheaterReferId uint32  `json:"theaterId" gorm:"primary_key;autoIncrement:false"`
	Theater        Theater `json:"-" gorm:"foreignKey:TheaterReferId;References:TheaterId"`
}
