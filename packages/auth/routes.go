package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()

	// Login Routes
	r.Get("/login", LoginHandler)
	r.Post("/login", LoginPostHandler)

	// Register Routes
	r.Get("/register", RegisterHandler)
	r.Post("/register", RegisterPostHandler)

	// Logout Routes
	r.Get("/logout", LogoutHandler)

	// Forgot Password Handling Routes
	r.Get("/forgot-password", ForgotPasswordHandler)

	// Debugging Purposes Only
	r.Get("/test-login", TestLoginHandler)
	return r
}
