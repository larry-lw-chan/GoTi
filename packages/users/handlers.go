package users

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Check if provided password matches
	render.Template(w, "/users/profile.app.tmpl", data)
}
