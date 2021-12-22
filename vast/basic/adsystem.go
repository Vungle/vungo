package vastbasic

// AdSystem type represents information about the ad server that returned the ad.
type AdSystem struct {
	System  string `xml:",chardata"`
	Version string `xml:"version,attr,omitempty"`
}

// Validate method validate AdSystem vast element
func (as *AdSystem) Validate(version Version) error {
	if len(as.System) == 0 {
		return ValidationError{Errs: []error{ErrAdSystemMissSystem}}
	}
	return nil
}
