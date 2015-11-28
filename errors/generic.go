package errors

import goErrors "errors"

//New returns a new generic error
func New(message string) error {
	return goErrors.New(message)
}
