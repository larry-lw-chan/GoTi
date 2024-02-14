package handlers

import (
	"net/http"

	"github.com/larry-lw-chan/go-tiny/packages/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "home.page.tmpl", nil)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "about.page.tmpl", nil)
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "contact.page.tmpl", nil)
}
