package pages

import (
	"net/http"

	"github.com/larry-lw-chan/goti/utils/render"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl", nil)
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl", nil)
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "contact.page.tmpl", nil)
}

func PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "privacy.page.tmpl", nil)
}
