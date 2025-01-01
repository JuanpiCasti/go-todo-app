package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/juanpicasti/go-todo-app/app/config"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() error {
	constr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s sslrootcert=%s",
		config.CFG.DatabaseUser,
		config.CFG.DatabasePassword,
		config.CFG.DatabaseHost,
		config.CFG.DatabasePort,
		config.CFG.DatabaseName,
		config.CFG.DatabaseSslMode,
		config.CFG.Sslrootcert,
	)

	dbConnection, err := sqlx.Connect("postgres", constr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	err = dbConnection.Ping()

	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	DB = dbConnection
	return err
}
