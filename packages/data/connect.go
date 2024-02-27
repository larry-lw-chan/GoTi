package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Allows the database to be accessed from any models
var db *sql.DB
var err error

func Connect() *sql.DB {
	db, err = sql.Open("sqlite3", "./data/sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	fmt.Println("Database Connected Successfully")
	return db
}
