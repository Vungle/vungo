package vast2

import vastbasic "github.com/Vungle/vungo/vast/basic"

// Version type represents the VAST document version.
type Version string

// Version2 is the supported version of this package.
const Version2 Version = "2.0"

// Validate method validates the Version according to the VAST.
func (v Version) Validate() error {
	if v != Version2 {
		return vastbasic.ValidationError{Errs: []error{vastbasic.ErrUnsupportedVersion}}
	}

	return nil
}
