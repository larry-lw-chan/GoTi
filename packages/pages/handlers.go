package pages

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/flash"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flash := flash.Get(w, r)
	if flash != nil {
		data["Flash"] = flash
	}
	render.Template(w, "/pages/home.base.tmpl", data)
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	render.Template(w, "/pages/about.base.tmpl", data)
}

func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	render.Template(w, "/pages/contact.base.tmpl", data)
}

func PrivacyPageHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	render.Template(w, "/pages/privacy.base.tmpl", data)
}
