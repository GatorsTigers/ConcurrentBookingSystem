package database

import (
	"fmt"
	"sync"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/logger"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"gorm.io/driver/mysql" // Or any other database driver you are using
	"gorm.io/gorm"
)

var (
	DbInstance *DatabaseInstance
	once       sync.Once
)

type DatabaseInstance struct {
	db *gorm.DB
}

func NewDatabaseClient(config *config.Config) *DatabaseInstance {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Dbname,
	)

	database, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		logger.Fatal("Could not connect to the database")
		return nil
	}

	logger.Info(fmt.Sprintf("Connected to the database %s", config.DB.Dbname))
	dbInstance := &DatabaseInstance{
		db: database,
	}
	return dbInstance
}

func InitDB(config *config.Config) {
	once.Do(func() {
		instance := NewDatabaseClient(config)
		DbInstance = instance
	})
}

func (i *DatabaseInstance) Ready() bool {
	var ready string
	txn := i.db.Raw("SELECT 1 as ready").Scan(&ready)
	if txn.Error != nil {
		return false
	} else if ready == "1" {
		return true
	}
	return false
}

func (i *DatabaseInstance) CreateTables() {
	if err := i.db.AutoMigrate(
		&models.City{},
		&models.Show{},
		&models.Theater{},
		&models.User{},
		&models.TheaterShow{},
		&models.Screen{},
		&models.Seat{},
		&models.ScreenShowSchedule{},
		&models.Ticket{},
	); err != nil {
		logger.Fatal("Failed to create the tables")
	}
}
