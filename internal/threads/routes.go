package threads

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

	// Thread Routes
	r.Get("/", IndexThreadHandler)
	r.Get("/new", NewThreadHandler)
	r.Post("/new", NewPostThreadHandler)
	r.Get("/show/{thread_id}", ShowThreadHandler)

	// Like Routes
	r.Post("/likes/{thread_id}", LikeThreadHandler)

	// Partial Routes for HTMX usage
	r.Get("/user-threads/{user_id}", UserThreadsHandler)
	r.Get("/user-replies/{user_id}", UserRepliesHandler)
	r.Get("/user-repost/{user_id}", UserRepostHandler)
	return r
}
