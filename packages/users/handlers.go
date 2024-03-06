package users

import (
	"context"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/data"
	"github.com/larry-lw-chan/goti/utils/flash"
	"github.com/larry-lw-chan/goti/utils/render"
)

// Authentication Handlers - TODO
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "forgot-password.page.tmpl", nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flash := flash.Get(w, r)
	if flash != nil {
		data["Flash"] = flash
	}
	render.Template(w, "register.page.tmpl", data)
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Handle Form Validation
	errs := validateCreateUser(r)
	if errs != nil {
		var message string
		for _, err := range errs {
			message += err.Error() + "<br />"
		}
		flash.Set(w, flash.FAILED, []byte(message))
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}

	// Generate Hashed Password
	hashPwd := HashPassword([]byte(r.FormValue("password")))

	// Insert new user into database
	queries := New(data.DB)
	ctx := context.Background()
	user := CreateUserParams{
		Username:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Password:  hashPwd,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	queries.CreateUser(ctx, user)

	// Todo - redirect to user authentication post
	flash.Set(w, flash.SUCCESS, []byte("Registration Worked!"))
	http.Redirect(w, r, "/register", http.StatusSeeOther)
}
