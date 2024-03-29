package pages

import (
	"errors"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/pages/home.base.tmpl")
}

func PagesHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get article slug and create template path
	pageSlug := chi.URLParam(r, "pageSlug")
	tmpl := "/pages/" + pageSlug + ".base.tmpl"

	// Check to see if template file exists
	if _, err := os.Stat(render.TmplPath + tmpl); errors.Is(err, os.ErrNotExist) {
		data["Code"] = http.StatusText(http.StatusNotFound)
		render.Template(w, data, "/pages/not_found.base.tmpl")
		return
	}

	// Render the template
	render.Template(w, data, tmpl)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	data["Code"] = http.StatusText(http.StatusNotFound)
	render.Template(w, data, "/pages/not_found.base.tmpl")
}
