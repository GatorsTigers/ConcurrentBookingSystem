package models

type Screen struct {
	ScreenId       uint32  `json:"screenId" gorm:"primary_key;"`
	ScreenName     string  `json:"screenName" gorm:"size:100;"`
	TheaterReferId uint32  `json:"theaterId"`
	Theater        Theater `json:"-" gorm:"foreignKey:TheaterReferId;References:TheaterId"`
}
