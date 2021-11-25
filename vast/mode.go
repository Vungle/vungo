package vast

// Mode type represents the various ways of how a companion ad can be displayed.
type Mode string

// Enumeration of possible modes of how a companion ad can be displayed.
const (
	ModeAll  Mode = "all"  // All companion ads must be displayed.
	ModeAny  Mode = "any"  // At least one companion ad must be displayed.
	ModeNone Mode = "none" // Companion ad display is optional.
)

// Validate method validates Mode according to the VAST.
func (mode Mode) Validate() error {
	if mode != ModeAll && mode != ModeAny && mode != ModeNone {
		return ValidationError{Errs: []error{ErrUnsupportedMode}}
	}

	return nil
}
