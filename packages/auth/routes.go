package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()
	r.Get("/login", LoginHandler)
	r.Get("/register", RegisterHandler)
	r.Post("/register", RegisterPostHandler)
	r.Get("/forgot-password", ForgotPasswordHandler)

	// Debugging Purposes Only
	r.Get("/test-auth", TestAuthHandler)
	r.Get("/secret", Secret)

	return r
}

// func Routes(r *chi.Mux) {
// 	r.Get("/login", LoginHandler)
// 	r.Get("/register", RegisterHandler)
// 	r.Post("/register", RegisterPostHandler)
// 	r.Get("/forgot-password", ForgotPasswordHandler)

// 	// Debugging Purposes Only
// 	r.Get("/test-auth", TestAuthHandler)
// 	r.Get("/secret", Secret)
// }
