package pages

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux) {
	r.Get("/", HomePageHandler)
	r.Get("/about", AboutPageHandler)
	r.Get("/contact", ContactPageHandler)
	r.Get("/privacy", PrivacyPageHandler)
}
