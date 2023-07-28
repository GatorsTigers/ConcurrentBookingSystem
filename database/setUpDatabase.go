package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"gorm.io/driver/mysql" // Or any other database driver you are using
	"gorm.io/gorm"
)

func InitializeDatabase() {
	dsn := "root:gatortiger@tcp(localhost:3306)/ConcurrentBookingSystem"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&models.City{}, &models.Show{}, &models.Theater{}, &models.User{}, &models.TheaterShow{}, &models.Screen{}, &models.Seat{},
		&models.ScreenShowSchedule{}, &models.Ticket{})

	if err != nil {
		panic("Failed to auto-migrate the table")
	}
}
