package database

import (
	"fmt"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/google/uuid"
	"gorm.io/driver/mysql" // Or any other database driver you are using
	"gorm.io/gorm"
)

type Database struct {
	dbInstance *gorm.DB
}

func (db *Database) InitializeDatabase(config *config.Config) {
	var database *gorm.DB
	var err error
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Dbname,
	)

	database, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to db")
	db.dbInstance = database
}

func (db *Database) CreateTables() {
	err := db.dbInstance.AutoMigrate(
		&models.City{},
		&models.Show{},
		&models.Theater{},
		&models.User{},
		&models.TheaterShow{},
		&models.Screen{},
		&models.Seat{},
		&models.ScreenShowSchedule{},
		&models.Ticket{},
	)

	if err != nil {
		panic("Failed to auto-migrate the table")
	}
}

func (db *Database) InsertDataIntoTables() {
	uuid := uuid.New()
	user := &models.User{
		EmailId: "jack@gmail.com",
		PhoneNo: "9900000000",
	}
	city := &models.City{
		CityId:   uuid.ID(),
		CityName: "NewYork",
	}
	db.dbInstance.Save(&user)
	db.dbInstance.Save(&city)
}
