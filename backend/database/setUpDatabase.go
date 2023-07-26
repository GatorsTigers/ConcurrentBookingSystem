package database

import (
	"github.com/GatorsTigers/ConcurrentBookingSystem/backend/models"
	"gorm.io/driver/mysql" // Or any other database driver you are using
	"gorm.io/gorm"
)

func initializeDatabase() {
	db, err := gorm.Open(mysql.Open("your-database-connection-string"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	err = db.AutoMigrate(&models.City{})
	if err != nil {
		panic("Failed to auto-migrate the table")
	}

}
