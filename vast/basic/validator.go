package vastbasic

// Validator is used as the validation interface
type Validator interface {
	Validate(version Version) error
}
