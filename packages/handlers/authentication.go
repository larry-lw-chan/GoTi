package handlers

import (
	"net/http"

	"github.com/larry-lw-chan/go-tiny/packages/render"
	"github.com/larry-lw-chan/go-tiny/packages/utils"
)

// Authentication Handlers - TODO
func Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	render.Html(w, "forgot-password.page.tmpl", nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flash := utils.GetFlash(w, r, utils.FlashFailed)
	if flash != nil {
		data["Flash"] = flash
		render.Html(w, "register.page.tmpl", data)
	} else {
		render.Html(w, "register.page.tmpl", nil)
	}
}

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	// Testing to see if the form is being submitted
	utils.SetFlash(w, utils.FlashFailed, []byte("Registration Failed!"))
	http.Redirect(w, r, "/register", http.StatusSeeOther)
}
