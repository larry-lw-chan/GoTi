package pages

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/internal/sessions/flash"
)

func Router() http.Handler {
	r := chi.NewRouter()

	// Load Flash Middleware
	r.Use(flash.CheckForFlash)

	// Login and Register Routes
	r.Get("/", HomePageHandler)
	r.Get("/pages/{pageSlug:[a-z-]+}", PagesHandler)

	return r
}
