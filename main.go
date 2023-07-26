package main

import (
	"fmt"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
)

func main() {
	database.InitializeDatabase()
	fmt.Print("CBS")
}
