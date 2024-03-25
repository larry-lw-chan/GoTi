package timelines

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/packages/auth"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
)

func Router() http.Handler {
	r := chi.NewRouter()

	// Load Middleware
	r.Use(flash.CheckForFlash)
	r.Use(auth.CheckIfAuthenticated)

	// Timeline Routes
	r.Get("/", TimelineIndexHandler)

	return r
}
