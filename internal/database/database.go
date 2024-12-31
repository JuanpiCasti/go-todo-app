package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	var err error
	constr := "user=root dbname=todo_app password=root host=localhost sslmode=disable"
	DB, err := sqlx.Connect("postgres", constr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	return DB, err
}
