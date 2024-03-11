package pages

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "pages/home.tmpl", nil)
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "pages/about.tmpl", nil)
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "pages/contact.tmpl", nil)
}

func PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "pages/privacy.tmpl", nil)
}
