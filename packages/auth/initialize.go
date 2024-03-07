package auth

import (
	"encoding/base64"
	"encoding/gob"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// Store and Session keys
const (
	COOKIE_NAME  = "cookie-store"
	USER_SESSION = "user-session"
)

// Session store
var store *sessions.CookieStore

func InitializeStore() {
	// Sets random key pair on startup if .env is not defined properly
	// Do NOT use this in production... set the .env file properly
	if os.Getenv("AUTH_KEY") == "" || os.Getenv("ENCRYPTION_KEY") == "" {
		store = sessions.NewCookieStore(
			securecookie.GenerateRandomKey(64),
			securecookie.GenerateRandomKey(32),
		)
	}

	// Decode the keys from base64 format
	authKey, _ := base64.StdEncoding.DecodeString(os.Getenv("AUTH_KEY"))
	encryptionKey, _ := base64.StdEncoding.DecodeString(os.Getenv("ENCRYPTION_KEY"))

	// Create a new cookie store
	store = sessions.NewCookieStore(
		[]byte(authKey),
		[]byte(encryptionKey),
	)

	// Set cookie store expiration time
	store.MaxAge(604800) // Default to 1 week

	// Set cookie store expiration time from .env
	if os.Getenv("MAX_AGE") != "" {
		maxAge, err := strconv.Atoi(os.Getenv("MAX_AGE"))
		if err != nil {
			log.Println("Error converting MAX_AGE to int")
		}
		store.MaxAge(maxAge)
	}

	// Needed to make work with gorilla/sessions
	gob.Register(UserSession{})
}
