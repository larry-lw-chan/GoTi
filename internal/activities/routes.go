package activities

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/internal/auth"
	"github.com/larry-lw-chan/goti/internal/sessions/flash"
)

// Protected Routes
func Router() http.Handler {
	r := chi.NewRouter()

	// Load Middleware
	r.Use(flash.CheckForFlash)
	r.Use(auth.CheckIfAuthenticated)

	// Relationship Routes
	r.Get("/", IndexActivityHandler)

	// Relationship Partials
	r.Get("/all", GetAllHandler)
	r.Get("/follows", GetFollowHandler)
	r.Get("/replies", GetRepliesHandler)
	r.Get("/repost", GetRepostHandler)

	return r
}
