package pages

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()

	// Login and Register Routes
	r.Get("/", HomePageHandler)

	r.Route("/pages", func(r chi.Router) {
		r.Get("/about", AboutPageHandler)
		r.Get("/contact", ContactPageHandler)
		r.Get("/privacy", PrivacyPageHandler)
	})

	return r
}
