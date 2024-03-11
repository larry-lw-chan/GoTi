package translate

import (
	"fmt"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	english locales.Translator      = en.New()
	uni     *ut.UniversalTranslator = ut.New(english, english)
)

// Errors translates validation errors into human friendly messages
func Errors(err error, validate *validator.Validate) (errs []error) {
	if err == nil {
		return nil
	}

	// Translate Error
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
