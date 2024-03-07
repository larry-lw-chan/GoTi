package users

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/cookie"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	user := cookie.GetUserSession(r)
	if auth := user.Authenticated; !auth {
		w.Write([]byte("You not authenticated"))
		return
	}
	w.Write([]byte("User Profile"))
}
