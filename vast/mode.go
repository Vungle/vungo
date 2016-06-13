package vast

// Mode type represents the various ways of how a companion ad can be displayed.
type Mode string

// Enumeration of possible modes of how a companion ad can be displayed.
const (
	MODE_ALL  Mode = "all"  // All companion ads must be displayed.
	MODE_ANY  Mode = "any"  // At least one companion ad must be displayed.
	MODE_NONE Mode = "none" // Companion ad display is optional.
)

// Validate method validates Mode according to the VAST.
func (mode Mode) Validate() error {

	if mode != MODE_ALL && mode != MODE_ANY && mode != MODE_NONE {
		return ErrUnsupportedMode
	}

	return nil
}
