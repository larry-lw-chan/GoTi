package main

import (
	"fmt"
	"net/http"

	"github.com/larry-lw-chan/go-tiny/packages/data"
	"github.com/larry-lw-chan/go-tiny/packages/router"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect Database
	db := data.Connect()
	defer db.Close() // Defer close the database connection

	// Get Routes
	r := router.Routes()

	// Start Server
	fmt.Println("Server is running at :3000")
	http.ListenAndServe(":3000", r)
}
