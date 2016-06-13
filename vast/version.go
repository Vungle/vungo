package vast

// Version type represents the VAST document version.
type Version string

// Enumeration of possible versions that package vast can support.
const (
	VERSION_2 Version = "2.0"
	VERSION_3 Version = "3.0"
)

// Validate method validates the Version according to the VAST.
func (v Version) Validate() error {

	if v != VERSION_2 && v != VERSION_3 {
		return ErrUnsupportedVersion
	}

	return nil
}
