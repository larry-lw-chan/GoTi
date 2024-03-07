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
	"golang.org/x/crypto/bcrypt"
)

// Store and Session keys
const (
	COOKIE_NAME  = "cookie-store"
	USER_SESSION = "user-session"
)

// Session store
var store *sessions.CookieStore

/*************************************************
* SESSION SERVICES
*************************************************/
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

// GetUserSession returns a user from session
// on error returns an empty user
func GetUserSession(r *http.Request) UserSession {
	session, _ := store.Get(r, COOKIE_NAME)
	value := session.Values[USER_SESSION]

	// Cast the value interface to UserSession
	userSession, ok := value.(UserSession) // Type assertion

	// If user not logged in
	if !ok {
		return UserSession{Authenticated: false}
	}

	return userSession // User is authenticated
}

// Delete a user authentication
func DelUserSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, COOKIE_NAME)
	session.Values["user"] = UserSession{}
	session.Options.MaxAge = -1
	session.Save(r, w)
}

/*************************************************
* HASHING SERVICES
*************************************************/
func HashPassword(pwd []byte) string {
	bytes, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(bytes)
}

func CheckPasswordHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

/*************************************************
* COOKIE STORE INITIALIZATION
*************************************************/
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
