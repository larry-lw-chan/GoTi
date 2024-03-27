package activities

import (
	"net/http"

	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func IndexActivityHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, nil, "/activities/index.app.tmpl")
}

func AllRelationsHandler(w http.ResponseWriter, r *http.Request) {
	// queries := New(database.DB)
	render.Template(w, nil, "/activities/all.app.tmpl")
}
