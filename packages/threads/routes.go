package threads

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/packages/auth"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
)

// Protected Routes
func Router() http.Handler {
	r := chi.NewRouter()

	// Load Middleware
	r.Use(flash.CheckForFlash)
	r.Use(auth.CheckIfAuthenticated)

	// Profile Routes
	r.Get("/new", NewThreadHandler)
	r.Post("/new", NewPostThreadHandler)

	// Partial Routes for HTMX usage
	r.Get("/user-threads", UserThreadsHandler)
	r.Get("/user-replies", UserRepliesHandler)
	r.Get("/user-repost", UserRepostHandler)
	return r
}
