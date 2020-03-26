package errs

import "errors"

// common errs
var (
	ErrBadRequest = errors.New("err bad request")
	ErrStorage    = errors.New("storage error")
)
