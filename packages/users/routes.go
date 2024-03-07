package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/profile", ProfileHandler)

	return r
}
