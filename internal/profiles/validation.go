package profiles

import (
	"github.com/go-playground/validator/v10"
	"github.com/larry-lw-chan/goti/internal/helper"
)

// use a single instance , it caches struct info
var (
	validate *validator.Validate
)

type CreateProfileValidation struct {
	Username string `validate:"required,min=4,max=15"`
	Name     string `validate:"omitempty,min=4,max=50"`
	Bio      string `validate:"omitempty,min=4,max=300"`
	Link     string `validate:"omitempty,url"`
	Avatar   string `validate:"omitempty,url"`
	Private  int64  `validate:"number"`
}

func validateCreateProfile(createProfileValidation CreateProfileValidation) (errs []error) {
	// Validate User Input
	validate = validator.New()

	vErrs := validate.Struct(&createProfileValidation)
	if vErrs != nil {
		errs = append(errs, helper.TranslateErrors(vErrs, validate)...)
	}

	// Return nil if no errors
	return errs
}
