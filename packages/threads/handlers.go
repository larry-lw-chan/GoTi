package threads

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/sessions/flash"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func NewThreadHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "/threads/new.app.tmpl", nil)
}

func NewPostThreadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Handle Form Validation
	errs := validateNewThread(r)
	if errs != nil {
		var message string
		for _, err := range errs {
			message += err.Error() + "<br />"
		}
		flash.Set(w, r, flash.ERROR, message)
		http.Redirect(w, r, "/threads/new", http.StatusSeeOther)
		return
	}
	render.Template(w, "/threads/new.app.tmpl", nil)
}

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
