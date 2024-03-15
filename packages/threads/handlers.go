package threads

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/utils/render"
)

/********************************************************
* Partials
*********************************************************/
func UserThreadsHandler(w http.ResponseWriter, r *http.Request) {
	render.Partial(w, "/threads/partials/userthreads.app.tmpl", nil)
}
