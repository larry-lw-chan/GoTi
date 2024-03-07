package auth

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	// Authenticated Routes Testing
	r.Get("/test-auth", TestAuthHandler)
	r.Get("/secret", Secret)
}
