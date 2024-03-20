package threads

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
	"github.com/larry-lw-chan/goti/packages/users"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func NewThreadHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]any)
	render.Template(w, "/threads/new.app.tmpl", data)
}

func NewPostThreadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Handle Form Validation
	errs := validateNewThread(r)
	if errs != nil {
		var message string
		for _, err := range errs {
			message += err.Error() + "\n"
		}
		flash.Set(w, r, flash.ERROR, message)
		http.Redirect(w, r, "/threads/new", http.StatusSeeOther)
		return
	}

	// Get user session information
	userSession := users.GetUserSession(r)

	// Create new thread
	queries := New(database.DB)
	threadParam := CreateThreadParams{
		Content:   r.FormValue("content"),
		ThreadID:  sql.NullInt64{Valid: false},
		UserID:    userSession.Id,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	_, err := queries.CreateThread(r.Context(), threadParam)
	if err != nil {
		flash.Set(w, r, flash.ERROR, "Failed to create new thread.  Please contact support.")
		http.Redirect(w, r, "/threads/new", http.StatusSeeOther)
		return
	}

	// Set flash message and render template
	flash.Set(w, r, flash.SUCCESS, "New thread created.")
	http.Redirect(w, r, "/threads/new", http.StatusSeeOther)
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
