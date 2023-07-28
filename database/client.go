package database

import (
	"fmt"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/google/uuid"
	"gorm.io/driver/mysql" // Or any other database driver you are using
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

func NewDatabaseClient(config *config.Config) (Client, error) {
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
	client := Client{
		DB: database,
	}
	return client, nil
}

func (c Client) Ready() bool {
	var ready string
	txn := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if txn.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

func (c Client) CreateTables() error {
	err := c.DB.AutoMigrate(
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
		return fmt.Errorf("failed to auto-migrate the table")
	}
	return nil
}

func (c Client) InsertDataIntoTables() error {
	uuid := uuid.New()
	user := &models.User{
		EmailId: "jack@gmail.com",
		PhoneNo: "9900000000",
	}
	city := &models.City{
		CityId:   uuid.ID(),
		CityName: "NewYork",
	}
	txn := c.DB.Save(&user)
	if txn.Error != nil {
		return txn.Error
	}
	txn = c.DB.Save(&city)
	if txn.Error != nil {
		return txn.Error
	}
	return nil
}
