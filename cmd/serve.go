package main

import (
	"github.com/juanpicasti/go-todo-app/config"
	"github.com/juanpicasti/go-todo-app/database"
	"github.com/juanpicasti/go-todo-app/router"
	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
)

func main() {
	// Load environment variables
	config.LoadConfig()
	// Initialize database connection
	db, err := database.Connect()
	if err != nil {
		log.Error().Err(err).Msg("Error connecting to database: ")
		panic(err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error closing database connection: ")
			panic(err)
		}
	}(db)

	// Initialize server
	r := router.SetupRouter(db)
	log.Info().Msg("Starting server on port " + config.CFG.ServerPort)
	err = r.Run(config.CFG.ServerPort)
	if err != nil {
		log.Error().Err(err).Msg("Error starting server: ")
		panic(err)
	}
}
