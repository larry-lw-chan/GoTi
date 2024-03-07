package auth

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/larry-lw-chan/goti/utils/translate"
)

// use a single instance , it caches struct info
var (
	validate *validator.Validate
)

type UserValidation struct {
	Username string `validate:"required,min=4,max=15"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func validateCreateUser(r *http.Request) (errs []error) {
	// Check if passwords match
	if r.FormValue("password") != r.FormValue("confirm_password") {
		errs = append(errs, errors.New("passwords do not match"))
	}

	// Check if privacy policy is agreed
	if r.FormValue("privacy") == "off" {
		errs = append(errs, errors.New("please agree to the privacy policy"))
	}

	// Validate User Input
	validate = validator.New()

	user := UserValidation{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	vErrs := validate.Struct(&user)
	if vErrs != nil {
		errs = append(errs, translate.Errors(vErrs, validate)...)
	}

	// Return nil if no errors
	return errs
}
