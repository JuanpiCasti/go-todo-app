package main

import (
	"github.com/juanpicasti/go-todo-app/internal/database"
	"github.com/juanpicasti/go-todo-app/internal/router"
)

func main() {

	// Initialize the database connection

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := router.SetupRouter(db)

	r.Run(":8080")

}
