package profiles

import (
	"context"
	"net/http"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/packages/sessions/cookie"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Find user by username
	userSession := cookie.GetUserSession(r)
	queries := New(database.DB)

	profile, err := queries.GetProfileFromUserId(context.Background(), 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data["Username"] = userSession.Username
	data["Profile"] = profile

	// Load user into profile
	render.Template(w, "/users/profile.app.tmpl", data)
}
