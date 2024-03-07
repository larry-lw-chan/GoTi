package pages

import (
	"net/http"

	"github.com/larry-lw-chan/goti/utils/render"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "page/home.html", nil)
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "page/about.html", nil)
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "page/contact.html", nil)
}

func PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "page/privacy.html", nil)
}
