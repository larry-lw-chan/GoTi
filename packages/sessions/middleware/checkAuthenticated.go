package middleware

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/sessions/cookie"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
)

func CheckIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := cookie.GetUserSession(r)
		if auth := user.Authenticated; !auth {
			flash.Set(w, r, flash.ERROR, "Please login to access.")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
