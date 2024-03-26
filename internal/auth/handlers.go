package auth

import (
	"net/http"

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
	luv := LoginUserValidation{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	errs := validateLoginUser(&luv)

	if errs != nil {
		messages := getErrorMessages(errs)
		flash.Set(w, r, flash.ERROR, messages)
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
	cuv := CreateUserValidation{
		Email:           r.FormValue("email"),
		Password:        r.FormValue("password"),
		ConfirmPassword: r.FormValue("confirm_password"),
		Privacy:         r.FormValue("privacy"),
	}

	errs := validateCreateUser(&cuv)

	if errs != nil {
		messages := getErrorMessages(errs)
		flash.Set(w, r, flash.ERROR, messages)
		http.Redirect(w, r, "/auth/register", http.StatusSeeOther)
		return
	}

	// Create new user
	user, err := CreateNewUser(r.FormValue("email"), r.FormValue("password"))

	// Check to make sure there is no existing user with that email
	if err != nil {
		flash.Set(w, r, flash.ERROR, "This email is already registered")
		http.Redirect(w, r, "/auth/register", http.StatusSeeOther)
		return
	}

	// Create user session and redirect to profile page
	CreateUserSession(w, r, user)
	flash.Set(w, r, flash.SUCCESS, "Registration Worked!")
	http.Redirect(w, r, "/profiles/edit", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	DeleteUserSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, nil, "/auth/forgot-password.base.tmpl")
}
