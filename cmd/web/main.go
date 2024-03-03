package main

import (
	"fmt"
	"net/http"

	"github.com/larry-lw-chan/go-tiny/database"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := database.Connect() // Connect Database
	defer db.Close()         // Defer close the database connection

	// Start Server
	r := Routes() // Get Routes
	fmt.Println("Server is running at :3000")
	http.ListenAndServe(":3000", r)
}
