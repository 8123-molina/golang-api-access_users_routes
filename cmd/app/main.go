package main

import (
	"log"

	"github.com/8123-molina/golang-api-access_users_routes/configs/database"
	"github.com/8123-molina/golang-api-access_users_routes/internal/router"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	router.Initialize()
}
