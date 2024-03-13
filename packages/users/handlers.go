package users

import (
	"net/http"

	"github.com/larry-lw-chan/goti/packages/sessions/cookie"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	if f := flash.Get(w, r); f != nil {
		data["Flash"] = f
	}

	user := cookie.GetUserSession(r)
	if auth := user.Authenticated; !auth {
		flash.Set(w, r, flash.ERROR, "Please login to view your profile.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Check if provided password matches
	render.Template(w, "/users/profile.app.tmpl", data)
}
