package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Authentication Routes
func AuthRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/login", LoginHandler)
	r.Post("/login", LoginPostHandler)
	r.Get("/register", RegisterHandler)
	r.Post("/register", RegisterPostHandler)
	r.Get("/logout", LogoutHandler)
	r.Get("/forgot-password", ForgotPasswordHandler)

	// Debugging Purposes Only
	r.Get("/test", TestLoginHandler)
	return r
}

// Protected Routes
func UserRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(CheckIfAuthenticated)
	r.Get("/profile", ProfileHandler)

	return r
}
