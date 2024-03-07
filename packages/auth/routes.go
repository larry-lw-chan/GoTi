package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()

	// Login and Register Routes
	r.Get("/login", LoginHandler)
	r.Post("/login", LoginPostHandler)
	r.Get("/register", RegisterHandler)
	r.Post("/register", RegisterPostHandler)

	// Forgot Password Handling Routes
	r.Get("/forgot-password", ForgotPasswordHandler)

	// Debugging Purposes Only
	r.Get("/test-auth", TestAuthHandler)
	r.Get("/secret", Secret)

	return r
}
