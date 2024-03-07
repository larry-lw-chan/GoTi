package users

import (
	"log"
	"net/http"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello")
	w.Write([]byte("User Profile"))
}
