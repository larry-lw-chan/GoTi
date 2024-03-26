package auth

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/larry-lw-chan/goti/internal/utils/translate"
)

/************************************************************
* Validations
***********************************************************/

// Create User Validation
type CreateUserValidation struct {
	Email           string `validate:"required,email"`
	Password        string `validate:"required,min=8"`
	ConfirmPassword string `validate:"required"`
	Privacy         string `validate:"required"`
}

func validateCreateUser(cuv *CreateUserValidation) (errs []error) {
	// Check if passwords match
	if cuv.Password != cuv.ConfirmPassword {
		errs = append(errs, errors.New("passwords do not match"))
	}

	// Check if privacy policy is agreed
	if cuv.Privacy == "off" {
		errs = append(errs, errors.New("please agree to the privacy policy"))
	}

	// Validate User Input
	validate := validator.New()
	createUser := CreateUserValidation{
		Email:    cuv.Email,
		Password: cuv.Password,
	}
	vErrs := validate.Struct(&createUser)
	if vErrs != nil {
		errs = append(errs, translate.Errors(vErrs, validate)...)
	}

	// Returns nil if no errors
	return errs
}

// Login Validation
type LoginUserValidation struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func validateLoginUser(luv *LoginUserValidation) (errs []error) {
	// Validate User Input
	validate := validator.New()

	loginUser := LoginUserValidation{
		Email:    luv.Email,
		Password: luv.Password,
	}

	vErrs := validate.Struct(&loginUser)
	if vErrs != nil {
		errs = append(errs, translate.Errors(vErrs, validate)...)
	}

	// Returns nil if no errors
	return errs
}
