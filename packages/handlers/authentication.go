package handlers

import (
	"net/http"

	"github.com/larry-lw-chan/go-tiny/packages/render"
)

// Authentication Handlers - TODO
func Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "register.page.tmpl", nil)
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "forgot-password.page.tmpl", nil)
}
