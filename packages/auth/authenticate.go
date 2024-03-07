package auth

import (
	"encoding/base64"
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/larry-lw-chan/goti/utils/debug"
)

// Struct to Hold Session
type UserSession struct {
	Username      string
	Authenticated bool
}

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

func CreateUserSession(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, COOKIE_NAME)

	// Create User Session and set it to the session
	userSession := &UserSession{
		Username:      "larry",
		Authenticated: true,
	}
	session.Values[USER_SESSION] = userSession

	// Save user session
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getUser returns a user from session s
// on error returns an empty user
func getUser(s *sessions.Session) UserSession {
	userSession := UserSession{}

	values := s.Values[USER_SESSION]
	debug.Print(values)
	// userSession, ok := values.(userSession)

	// // If user not authenticated
	// if !ok {
	// 	return UserSession{Authenticated: false}
	// }

	// // User is authenticated
	return userSession
}
