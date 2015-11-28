package rest

import (
	"fmt"
	"net/http"

	"github.com/tehleach/hue/errors"
)

//MatchErrorToCode matches error type to http status code
func MatchErrorToCode(err error) int {
	switch err.(type) {
	case *errors.NotFound:
		return http.StatusNotFound
	case *errors.OutOfBounds:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

//WriteError writes an error message to w
func WriteError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}
