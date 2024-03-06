package data

import (
	"database/sql"
	"fmt"
	"log"
)

// Global Database Connection
var DB *sql.DB

func Connect() *sql.DB {
	// Connect Database
	db, err := sql.Open("sqlite3", "./data/sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}

	// Test Database Connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}
	fmt.Println("Database Connected Successfully")

	// Inject Database to shared variable
	inject(db)
	return db
}

func inject(db *sql.DB) {
	DB = db
}
