package main

import (
	"fmt"

	db "github.com/Coluding/udemy_backend/db/sqlc"
)

func main() {
	// Get the database connection from the db package
	fmt.Println("szarting")
	database := db.GetDB()
	db.RunTestQuery(database)

	// Perform database operations here
	// For example, you can query the database or execute other operations

	fmt.Println("Do something with the database...")
}
