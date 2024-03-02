package handlers

import (
	"net/http"

	"github.com/larry-lw-chan/go-tiny/packages/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl", nil)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl", nil)
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "contact.page.tmpl", nil)
}

func PrivacyPage(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "privacy.page.tmpl", nil)
}
