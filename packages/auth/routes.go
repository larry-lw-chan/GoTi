package auth

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	r.Get("/login", LoginHandler)
	r.Get("/register", RegisterHandler)
	r.Post("/register", RegisterPostHandler)
	r.Get("/forgot-password", ForgotPasswordHandler)

	// Debugging Purposes Only
	r.Get("/test-auth", TestAuthHandler)
	r.Get("/secret", Secret)
}
