package threads

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/utils/render"
)

/********************************************************
* Partials
*********************************************************/
func UserThreadsHandler(w http.ResponseWriter, r *http.Request) {
	render.Partial(w, "/threads/partials/user_threads.app.tmpl", nil)
}

func UserRepliesHandler(w http.ResponseWriter, r *http.Request) {
	render.Partial(w, "/threads/partials/user_replies.app.tmpl", nil)
}

func UserRepostHandler(w http.ResponseWriter, r *http.Request) {
	render.Partial(w, "/threads/partials/user_repost.app.tmpl", nil)
}
