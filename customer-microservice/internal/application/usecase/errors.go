package usecase

import "errors"

var (
	errEmptyFields    error = errors.New("error: empty value for fied")
	errTooLong        error = errors.New("error: too long value for fied")
	errTooShort       error = errors.New("error: too short value for fied")
	errInvalidGenda   error = errors.New("error: invalid input genda")
	errInvalidId      error = errors.New("error: provided is invalid")
	errInvalidPhone   error = errors.New("error: invalid input phone number")
	errInvalidZipCode error = errors.New("error: invalid zip code")

	errSendObject     error = errors.New("error of sending bs object")
	errRetrieveObject error = errors.New("error: impossible to retrieve demanded objet")
)
