package profiles

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

	// Profile Routes
	r.Get("/show", ShowHandler)
	r.Get("/edit", EditHandler)
	r.Post("/edit", EditPostHandler)
	r.Get("/edit/photo", EditPhotoHandler)
	r.Post("/edit/photo", EditPhotoPostHandler)
	r.Post("/delete/photo", DeletePhotoPostHandler)

	// Show Other User Profiles Routes
	r.Get("/{username}", ShowProfileHandler)

	return r
}
