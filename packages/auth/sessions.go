package auth

import (
	"encoding/gob"
	"net/http"

	"github.com/larry-lw-chan/goti/packages/sessions/cookie"
)

// Auth Session Value key
const USER_SESSION string = "user-session"

// User Session Struct
type UserSession struct {
	ID            int64
	Uuid          string
	Username      string
	Authenticated bool
}

// Needed to make work with gorilla/sessions
func init() {
	gob.Register(UserSession{})
}

/*************************************************
* USER SESSION SERVICES
*************************************************/
func CreateUserSession(w http.ResponseWriter, r *http.Request, user *User) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := cookie.Store.Get(r, cookie.STORE)

	// Create User Session and set it to the session
	userSession := &UserSession{
		ID:            user.ID,
		Uuid:          user.Uuid,
		Username:      user.Username,
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
	session, _ := cookie.Store.Get(r, cookie.STORE)
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
func DeleteUserSession(w http.ResponseWriter, r *http.Request) {
	session, _ := cookie.Store.Get(r, cookie.STORE)
	session.Values["user"] = UserSession{}
	session.Options.MaxAge = -1
	session.Save(r, w)
}
