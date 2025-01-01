package main

import (
	"github.com/juanpicasti/go-todo-app/config"
	"github.com/juanpicasti/go-todo-app/database"
	"github.com/juanpicasti/go-todo-app/router"
	"log"

	"github.com/jmoiron/sqlx"
)

func main() {

	// Load environment variables
	cfg := config.LoadConfig()

	// Initialize database connection
	err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database connection: ", err)
		}
	}(database.DB)

	// Initialize server
	r := router.SetupRouter()
	err = r.Run(cfg.ServerPort)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
