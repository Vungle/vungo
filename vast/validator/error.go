package validator

import "fmt"

// FieldError is returned when a validation error occurs on the Field of an
// object with a given Reason.
type FieldError struct {
	Field  string
	Reason string
}

// Error implements the error interface.
func (v FieldError) Error() string {
	return fmt.Sprintf("%s %s", v.Field, v.Reason)
}

// ValidationError is returned when a validation error occurs on an object
// because of one or more FieldErrors.
type ValidationError struct {
	Errs []error
}

// Error implements the error interface.
func (v ValidationError) Error() string {
	return fmt.Sprintf("%+v", v.Errs)
}
