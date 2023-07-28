package main

import (
	"fmt"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	_db "github.com/GatorsTigers/ConcurrentBookingSystem/database"
)

func main() {
	config := config.GetConfig()
	db := _db.Database{}
	db.InitializeDatabase(config)
	db.CreateTables()
	db.InsertDataIntoTables()
	fmt.Printf("Successfullty Inserted")
}
