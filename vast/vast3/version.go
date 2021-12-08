package vast3

import vastbasic "github.com/Vungle/vungo/vast/basic"

// Version type represents the VAST document version.
type Version string

// Version3 is the supported version of this package.
const Version3 Version = "3.0"

// Validate method validates the Version according to the VAST.
func (v Version) Validate() error {
	if v != Version3 {
		return vastbasic.ValidationError{Errs: []error{vastbasic.ErrUnsupportedVersion}}
	}

	return nil
}
