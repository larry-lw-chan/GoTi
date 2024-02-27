package models

import (
	"fmt"
	"log"
)

func Test() {
	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	fmt.Println("Database Connected Successfully")
}
