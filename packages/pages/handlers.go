package pages

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "pages/home.html", nil)
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "pages/about.html", nil)
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "pages/contact.html", nil)
}

func PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "pages/privacy.html", nil)
}
