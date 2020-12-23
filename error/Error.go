package errors

import (
	"net/http"
)

// ValidationError used for Validation error
type ValidationError struct {
	ErrorCode int
	ErrorMsg  string
}

func (vd ValidationError) Error() string {
	return vd.ErrorMsg
}

// NewValidationError used to create new Validation error
func NewValidationError(msg string) ValidationError {
	return ValidationError{
		ErrorCode: http.StatusBadGateway,
		ErrorMsg:  msg,
	}
}
