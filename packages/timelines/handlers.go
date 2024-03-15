package timelines

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func TimelineIndexHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "/timelines/index.app.tmpl", nil)
}
