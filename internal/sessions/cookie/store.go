package cookie

import (
	"encoding/base64"
	"log"
	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// Options for the cookie store
type Options struct {
	AuthKey       string
	EncryptionKey string
	MaxAge        string
}

// Defaults for the cookie store
const (
	MAX_AGE int    = 604800
	STORE   string = "cookie-store"
)

// Store is the cookie store
var Store *sessions.CookieStore

func NewStore(o Options) {
	// Sets random key pair on startup if .env is not defined properly
	// Do NOT use this in production... set the .env file properly
	if o.AuthKey == "" || o.EncryptionKey == "" {
		authKey := securecookie.GenerateRandomKey(64)
		encryptionKey := securecookie.GenerateRandomKey(32)
		Store = sessions.NewCookieStore(authKey, encryptionKey)
	} else {
		// Decode the keys from base64 format
		authKey, _ := base64.StdEncoding.DecodeString(o.AuthKey)
		encryptionKey, _ := base64.StdEncoding.DecodeString(o.EncryptionKey)
		Store = sessions.NewCookieStore(authKey, encryptionKey)
	}

	// Set cookie store expiration time from .env
	if o.MaxAge != "" {
		maxAge, err := strconv.Atoi(o.MaxAge)
		if err != nil {
			log.Println("Error converting MAX_AGE to int")
		}
		Store.MaxAge(maxAge)
	} else {
		// Set cookie store expiration time
		Store.MaxAge(MAX_AGE) // Default to 1 week
	}
}
