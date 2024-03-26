package helper

// getErrorMessages convert validation errors into a string of error messages
func GetErrorMessages(errs []error) (message string) {
	for _, err := range errs {
		message += err.Error() + "<br />"
	}
	return message
}
