package vast2

// AdSystem type represents information about the ad server that returned the ad.
type AdSystem struct {
	Version string `xml:"version,attr,omitempty"`

	System string `xml:",chardata"`
}

// Validate methods validate the AdSystem element according to the VAST.
// System is required.
func (as *AdSystem) Validate() error {
	if len(as.System) == 0 {
		return ValidationError{Errs: []error{ErrAdSystemMissSystem}}
	}
	return nil
}
