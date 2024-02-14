package handlers

import (
	"net/http"

	"github.com/larry-lw-chan/go-tiny/packages/render"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "/pages/home.page.tmpl", nil)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "/pages/about.page.tmpl", nil)
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "/pages/contact.page.tmpl", nil)
}
