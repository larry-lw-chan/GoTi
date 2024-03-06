package users

import (
	"errors"
	"net/http"
)

func validateCreateUser(r *http.Request) (err error) {
	// Check if passwords match
	if r.FormValue("password") != r.FormValue("confirm_password") {
		return errors.New("passwords do not match")
	}

	// Check if privacy policy is agreed
	if r.FormValue("privacy") == "off" {
		return errors.New("please agree to the privacy policy")
	}

	// Validate User Input

	// Return nil if no errors
	return nil
}
