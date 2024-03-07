package auth

import (
	"net/http"
)

func TestAuthHandler(w http.ResponseWriter, r *http.Request) {
	CreateUserSession(w, r)
	w.Write([]byte("Test Auth Handler"))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, COOKIE_NAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := getUser(session)

	// if auth := user.Authenticated; !auth {
	// 	session.AddFlash("You don't have access!")
	// 	err = session.Save(r, w)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	http.Redirect(w, r, "/forbidden", http.StatusFound)
	// 	return
	// }

	w.Write([]byte("Secret working" + user.Username))
}
