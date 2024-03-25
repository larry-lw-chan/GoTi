package profiles

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/larry-lw-chan/goti/internal/utils/translate"
)

// use a single instance , it caches struct info
var (
	validate *validator.Validate
)

type CreateProfileValidation struct {
	Username string `validate:"required,min=4,max=15"`
	Name     string `validate:"omitempty,min=4,max=50"`
	Bio      string `validate:"omitempty,min=4,max=250"`
	Link     string `validate:"omitempty,url"`
	Avatar   string `validate:"omitempty,url"`
	Private  int64  `validate:"number"`
}

func validateCreateProfile(r *http.Request, private int64) (errs []error) {
	// Validate User Input
	validate = validator.New()

	createProfileValidation := CreateProfileValidation{
		Username: r.FormValue("username"),
		Name:     r.FormValue("name"),
		Bio:      r.FormValue("bio"),
		Link:     r.FormValue("link"),
		Avatar:   r.FormValue("avatar"),
		Private:  private,
	}
	vErrs := validate.Struct(&createProfileValidation)
	if vErrs != nil {
		errs = append(errs, translate.Errors(vErrs, validate)...)
	}

	// Return nil if no errors
	return errs
}
