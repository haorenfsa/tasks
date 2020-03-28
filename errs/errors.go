package errs

import (
	"errors"
	"net/http"
)

// common errs
var (
	ErrBadRequest = errors.New("err bad request")
	ErrStorage    = errors.New("storage error")
)

// ErrorToHTTPCode ..
func ErrorToHTTPCode(err error) int {
	switch err {
	case ErrBadRequest:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
