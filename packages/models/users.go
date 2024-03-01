package models

import (
	"fmt"
	"log"
)

type User struct {
	Username string
	Email    string
	Password string
}

func Test() {
	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	fmt.Println("Database Connected Successfully")
}
