package users

import (
	"fmt"
	"log"

	"github.com/larry-lw-chan/go-tiny/database"
)

type User struct {
	Username string
	Email    string
	Password string
}

func Test() {
	db := database.DB

	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot ping database because %s", err)
	}

	fmt.Println("Database Connected Successfully")
}
