package timelines

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
	"github.com/larry-lw-chan/goti/packages/sessions/middleware"
)

func Router() http.Handler {
	r := chi.NewRouter()

	// Load Middleware
	r.Use(flash.TryGetFlash)
	r.Use(middleware.CheckIfAuthenticated)

	// Timeline Routes
	r.Get("/", TimelineIndexHandler)

	return r
}
