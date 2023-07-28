package main

import (
	"log"

	"github.com/GatorsTigers/ConcurrentBookingSystem/config"
	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
)

func main() {
	config := config.GetConfig()
	db, err := database.NewDatabaseClient(config)
	if err != nil {
		log.Fatalf("failed to initialize daatabase Client: %s", err)
	}

	db.CreateTables()
	db.InsertDataIntoTables()
	log.Printf("Successfullty Inserted")
}
