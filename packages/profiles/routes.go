package profiles

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/packages/sessions/middleware"
)

// Protected Routes
func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.CheckIfAuthenticated)
	r.Get("/", ProfileHandler)

	return r
}
