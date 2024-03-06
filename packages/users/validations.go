package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
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
	validate := validator.New()

	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	user := UserValidation{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}
	validateErrs := validate.Struct(&user)

	if validateErrs != nil {
		validateErrs := translateError(validateErrs, trans)
		errs = append(errs, validateErrs...)
	}

	// Return nil if no errors
	return errs
}

// Makes Validation more human readable
func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
