package auth

import (
	"net/http"
)

// Get a session. We're ignoring the error resulted from decoding an
// existing session: Get() always returns a session, even if empty.
func CreateUserSession(w http.ResponseWriter, r *http.Request) {
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
