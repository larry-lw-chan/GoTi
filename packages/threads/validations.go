package threads

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/larry-lw-chan/goti/packages/utils/translate"
)

// use a single instance , it caches struct info
var (
	validate *validator.Validate
)

type NewThreadValidation struct {
	Content string `validate:"required,min=4,max=15"`
}

func validateNewThread(r *http.Request) (errs []error) {
	// Validate User Input
	validate = validator.New()

	newThread := NewThreadValidation{
		Content: r.FormValue("content"),
	}

	vErrs := validate.Struct(&newThread)
	if vErrs != nil {
		errs = append(errs, translate.Errors(vErrs, validate)...)
	}

	// Return nil if no errors
	return errs
}
