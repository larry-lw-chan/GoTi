package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
)

// Authentication Routes
func Router() http.Handler {
	r := chi.NewRouter()

	// Load Flash Middleware
	r.Use(flash.CheckForFlash)

	// Authentication Routes
	r.Get("/login", LoginHandler)
	r.Post("/login", LoginPostHandler)
	r.Get("/register", RegisterHandler)
	r.Post("/register", RegisterPostHandler)
	r.Get("/logout", LogoutHandler)
	r.Get("/forgot-password", ForgotPasswordHandler)
	return r
}
