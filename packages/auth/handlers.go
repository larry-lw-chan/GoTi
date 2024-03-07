package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/data"
	"github.com/larry-lw-chan/goti/packages/cookie"
	"github.com/larry-lw-chan/goti/packages/users"
	"github.com/larry-lw-chan/goti/utils/flash"
	"github.com/larry-lw-chan/goti/utils/render"
)

// Authentication Handlers - TODO
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "auth/login.html", nil)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "auth/forgot-password.html", nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flash := flash.Get(w, r)
	if flash != nil {
		data["Flash"] = flash
	}
	render.Template(w, "auth/register.html", data)
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
		http.Redirect(w, r, "/auth/register", http.StatusSeeOther)
	}

	// Generate Hashed Password
	hashPwd := HashPassword([]byte(r.FormValue("password")))

	// Insert new user into database
	queries := users.New(data.DB)
	ctx := context.Background()
	user := users.CreateUserParams{
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

func TestLoginHandler(w http.ResponseWriter, r *http.Request) {
	cookie.CreateUserSession(w, r)
	w.Write([]byte("Create User Session"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie.DeleteUserSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := cookie.GetUserSession(r)

	// Check if user is authenticated
	if auth := user.Authenticated; !auth {
		w.Write([]byte("You not authenticated"))
		return
	}

	// if auth := user.Authenticated; !auth {
	// 	session.AddFlash("You don't have access!")
	// 	err = session.Save(r, w)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	http.Redirect(w, r, "/forbidden", http.StatusFound)
	// 	return
	// }

	w.Write([]byte("Secret working " + user.Username))
}
