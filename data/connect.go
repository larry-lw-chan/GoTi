package data

import (
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB // Global Database Connection

// Database Connection Code
func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	fmt.Println("Database Connected Successfully")
	return db
}

func Inject(db *sql.DB) {
	DB = db
}
