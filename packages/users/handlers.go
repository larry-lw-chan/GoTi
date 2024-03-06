package users

import (
	"context"
	"net/http"

	"github.com/larry-lw-chan/goti/data"
	"github.com/larry-lw-chan/goti/utils/debug"
	"github.com/larry-lw-chan/goti/utils/flash"
	"github.com/larry-lw-chan/goti/utils/render"
)

// Authentication Handlers - TODO
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	queries := New(data.DB)
	ctx := context.Background()

	users, _ := queries.GetUsers(ctx)
	debug.Print(users)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "forgot-password.page.tmpl", nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flash := flash.Get(w, r, flash.FAILED)
	if flash != nil {
		data["Flash"] = flash
		render.Template(w, "register.page.tmpl", data)
	} else {
		render.Template(w, "register.page.tmpl", nil)
	}
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Testing to see if the form is being submitted
	flash.Set(w, flash.FAILED, []byte("Registration Failed!"))
	http.Redirect(w, r, "/register", http.StatusSeeOther)
}
