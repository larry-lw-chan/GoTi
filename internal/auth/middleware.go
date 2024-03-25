package auth

import (
	"net/http"

	"github.com/larry-lw-chan/goti/internal/sessions/flash"
)

func CheckIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := GetUserSession(r)

		if auth := user.Authenticated; !auth {
			flash.Set(w, r, flash.ERROR, "Please login to access.")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
