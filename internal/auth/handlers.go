package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/internal/sessions/flash"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

/*************************************************
* Authenticate Handlers
*************************************************/
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/auth/login.base.tmpl")
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Handle Form Validation
	errs := validateLoginUser(r)
	if errs != nil {
		var message string
		for _, err := range errs {
			message += err.Error() + ". "
		}
		flash.Set(w, r, flash.ERROR, message)
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		return
	}

	// Find user by email
	user, isAuthenticated := AuthenticateUser(r.FormValue("email"), r.FormValue("password"))

	// Check if provided password matches
	if isAuthenticated {
		CreateUserSession(w, r, user)
		flash.Set(w, r, flash.SUCCESS, "You have successfully logged in!")
		http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
	} else {
		flash.Set(w, r, flash.ERROR, "User not found or password incorrect")
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/auth/register.base.tmpl")
}

func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Handle Form Validation
	errs := validateCreateUser(r)
	if errs != nil {
		var message string
		for _, err := range errs {
			message += err.Error() + ".  "
		}
		flash.Set(w, r, flash.ERROR, message)
		http.Redirect(w, r, "/auth/register", http.StatusSeeOther)
	}

	// Generate Hashed Password
	hashPwd := HashPassword([]byte(r.FormValue("password")))

	// Insert new user into database
	queries := New(database.DB)
	createUser := CreateUserParams{
		Uuid:      generateUUID(),
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
		CreateUserSession(w, r, user)
		http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
	} else {
		flash.Set(w, r, flash.ERROR, "User not found or password incorrect")
		http.Redirect(w, r, "/auth/register", http.StatusSeeOther)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	DeleteUserSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, nil, "/auth/forgot-password.base.tmpl")
}
