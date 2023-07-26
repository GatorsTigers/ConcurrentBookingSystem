package db

import (
	"gorm.io/driver/mysql" // Or any other database driver you are using
	"gorm.io/gorm"
)

func initializeDatabase() {
	_, err := gorm.Open(mysql.Open("your-database-connection-string"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}
