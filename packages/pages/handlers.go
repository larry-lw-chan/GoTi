package pages

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/pages/home.base.tmpl")
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/pages/about.base.tmpl")
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/pages/contact.base.tmpl")
}

func PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/pages/privacy.base.tmpl")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	data["Code"] = http.StatusText(http.StatusNotFound)
	render.Template(w, data, "/pages/not_found.base.tmpl")
}
