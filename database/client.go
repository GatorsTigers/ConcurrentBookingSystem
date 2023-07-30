package database

import (
	"fmt"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"gorm.io/driver/mysql" // Or any other database driver you are using
	"gorm.io/gorm"
)

var DbInstance *DBInstance

// DBInstance represents the singleton database instance.
type DBInstance struct {
	db *gorm.DB
}

func NewDatabaseClient(config *config.Config) (*DBInstance, error) {
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
	client := &DBInstance{
		db: database,
	}
	return client, nil
}

func InitDB(config *config.Config) {
	instance, err := NewDatabaseClient(config)
	if err != nil {
		panic(err)
	}
	DbInstance = instance
}

func (c *DBInstance) GetDB() *gorm.DB {
	return c.db
}

func (c *DBInstance) Ready() bool {
	var ready string
	txn := c.db.Raw("SELECT 1 as ready").Scan(&ready)
	if txn.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

func CreateTables(db *gorm.DB) error {
	err := db.AutoMigrate(
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

func InsertDataIntoTables(db *gorm.DB) error {
	user := &models.User{
		EmailId: "jack@gmail.com",
		PhoneNo: "9900000000",
	}
	city := &models.City{
		CityName: "NewYork",
	}
	txn := db.Save(&user)
	if txn.Error != nil {
		return txn.Error
	}
	txn = db.Save(&city)
	if txn.Error != nil {
		return txn.Error
	}
	return nil
}
