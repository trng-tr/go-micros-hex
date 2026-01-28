package usecase

import "errors"

var (
	errInvalidVaue error = errors.New("error: invalid value for")
	errOccurred    error = errors.New("an error has occurred")
	errNotEnough   error = errors.New("error: product quantity not enough")
	errNoRows      error = errors.New("error: no rows in database")
)
