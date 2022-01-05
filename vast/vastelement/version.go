package vastelement

// Version type represents the VAST document version.
type Version string

// Version is the supported version of this package.
const (
	Version2 Version = "2.0"
	Version3 Version = "3.0"
	Version4 Version = "4.0"
)

// Validate method validates the Version according to the VAST.
func (v Version) Validate(version Version) error {
	if v != Version2 && v != Version3 && v != Version4 {
		return ValidationError{Errs: []error{ErrUnsupportedVersion}}
	}
	return nil
}
