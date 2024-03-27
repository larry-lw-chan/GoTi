package search

import (
	"net/http"

	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func IndexThreadHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, nil, "/search/index.app.tmpl")
}
