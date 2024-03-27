package helper

import (
	"net/http"

	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func PageNotFound(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		render.Template(w, nil, "/errors/404.app.tmpl")
	}
}
