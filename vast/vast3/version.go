package vast3

// Version type represents the VAST document version.
type Version string

// Enumeration of possible versions that package vast can support.
const (
	Version2 Version = "2.0"
	Version3 Version = "3.0"
)

// Validate method validates the Version according to the VAST.
func (v Version) Validate() error {
	if v != Version2 && v != Version3 {
		return ValidationError{Errs: []error{ErrUnsupportedVersion}}
	}

	return nil
}
