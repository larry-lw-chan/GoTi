package handlers

import (
	"net/http"

	"github.com/larry-lw-chan/go-tiny/utils"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "/pages/home.page.tmpl", nil)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "/pages/about.page.tmpl", nil)
}
