package database

import (
	"fmt"
	"github.com/juanpicasti/go-todo-app/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect(cfg *config.Config) error {
	constr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s sslrootcert=%s",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
		cfg.DatabaseSslMode,
		cfg.Sslrootcert,
	)

	dbConnection, err := sqlx.Connect("postgres", constr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return err
	}

	err = dbConnection.Ping()

	if err != nil {
		log.Fatal("Error pinging database: ", err)
		return err
	}

	DB = dbConnection

	return nil
}
