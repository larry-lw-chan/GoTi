package auth

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/data"
	"github.com/larry-lw-chan/goti/packages/cookie"
	"github.com/larry-lw-chan/goti/packages/flash"
	"github.com/larry-lw-chan/goti/packages/users"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

// Authentication Handlers - TODO
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flash := flash.Get(w, r)
	if flash != nil {
		data["Flash"] = flash
	}
	render.Template(w, "auth/login.tmpl", data)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Handle Form Validation
	errs := validateLoginUser(r)
	if errs != nil {
		var message string
		for _, err := range errs {
			message += err.Error() + "<br />"
		}
		flash.Set(w, r, flash.ERROR, message)
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	// Find user by email
	user, isAuthenticated := AuthenticateUser(r.FormValue("email"), r.FormValue("password"))

	// Check if provided password matches
	if isAuthenticated {
		cookie.CreateUserSession(w, r, user.Username)
		http.Redirect(w, r, "/users/profile", http.StatusSeeOther)
	} else {
		flash.Set(w, r, flash.ERROR, "User not found or password incorrect")
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
	}
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "auth/forgot-password.tmpl", nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flash := flash.Get(w, r)
	if flash != nil {
		data["Flash"] = flash
	}
	render.Template(w, "auth/register.tmpl", data)
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
		flash.Set(w, r, flash.ERROR, message)
		http.Redirect(w, r, "/auth/register", http.StatusSeeOther)
	}

	// Generate Hashed Password
	hashPwd := HashPassword([]byte(r.FormValue("password")))

	// Insert new user into database
	queries := users.New(data.DB)
	createUser := users.CreateUserParams{
		Username:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Password:  hashPwd,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	queries.CreateUser(context.Background(), createUser)

	// Authenticate user
	user, isAuthenticated := AuthenticateUser(r.FormValue("email"), r.FormValue("password"))

	// Check if provided password matches
	if isAuthenticated {
		flash.Set(w, r, flash.SUCCESS, "Registration Worked!")
		cookie.CreateUserSession(w, r, user.Username)
		http.Redirect(w, r, "/users/profile/", http.StatusSeeOther)
	} else {
		flash.Set(w, r, flash.ERROR, "User not found or password incorrect")
		http.Redirect(w, r, "/auth/register", http.StatusSeeOther)
	}
}

func TestLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(HashPassword([]byte("password")))

	// Create User Session
	cookie.CreateUserSession(w, r, "Testy")
	flash.Set(w, r, flash.SUCCESS, "User Session Created")
	w.Write([]byte("Create User Session"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie.DeleteUserSession(w, r)
	flash.Set(w, r, flash.SUCCESS, "You are successfully logged out")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
