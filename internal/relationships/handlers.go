package relationships

import (
	"net/http"

	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func IndexRelationHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, nil, "/relationships/index.app.tmpl")
}

func AllRelationHandler(w http.ResponseWriter, r *http.Request) {
	// queries := New(database.DB)
	render.Template(w, nil, "/relationships/all.app.tmpl")
}
