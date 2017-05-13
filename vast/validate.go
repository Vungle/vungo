package vast

import "fmt"

// Validator is the interface that wraps the Validate method which should be
// implemented by an object that can be validated.
type Validator interface {
	Validate() error
}

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
