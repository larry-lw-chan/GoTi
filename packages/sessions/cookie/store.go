package cookie

import (
	"encoding/base64"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// Store and Session keys
const (
	STORE = "cookie-store"
)

// Session store
var Store *sessions.CookieStore

/*************************************************
* COOKIE STORE INITIALIZATION
*************************************************/
func InitializeStore() {
	// Sets random key pair on startup if .env is not defined properly
	// Do NOT use this in production... set the .env file properly
	if os.Getenv("AUTH_KEY") == "" || os.Getenv("ENCRYPTION_KEY") == "" {
		Store = sessions.NewCookieStore(
			securecookie.GenerateRandomKey(64),
			securecookie.GenerateRandomKey(32),
		)
	}

	// Decode the keys from base64 format
	authKey, _ := base64.StdEncoding.DecodeString(os.Getenv("AUTH_KEY"))
	encryptionKey, _ := base64.StdEncoding.DecodeString(os.Getenv("ENCRYPTION_KEY"))

	// Create a new cookie store
	Store = sessions.NewCookieStore(
		[]byte(authKey),
		[]byte(encryptionKey),
	)

	// Set cookie store expiration time
	Store.MaxAge(604800) // Default to 1 week

	// Set cookie store expiration time from .env
	if os.Getenv("MAX_AGE") != "" {
		maxAge, err := strconv.Atoi(os.Getenv("MAX_AGE"))
		if err != nil {
			log.Println("Error converting MAX_AGE to int")
		}
		Store.MaxAge(maxAge)
	}
}
