package main

import (
	"fmt"
	"net/http"

	"github.com/larry-lw-chan/goti/data"
	"github.com/larry-lw-chan/goti/packages/routes"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := data.Connect() // Connect Database
	defer db.Close()     // Defer close the database connection
	data.Inject(db)      // Inject Database Connection

	// Start Server
	r := routes.Routes() // Get Routes
	fmt.Println("Server is running at :3000")
	http.ListenAndServe(":3000", r)
}
