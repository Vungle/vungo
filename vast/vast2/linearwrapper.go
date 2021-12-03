package vast2

import "github.com/Vungle/vungo/vast/basic"

// LinearWrapper type represents a <Linear> element within a <Wrapper> element that defines a
// wrapped linear creative.
type LinearWrapper struct {
	Trackings   []*vastbasic.Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks *vastbasic.VideoClicks `xml:"VideoClicks,omitempty"` // VideoClicks here is different from the Linear one.
}

// Validate method validates the LinearWrapper according to the VAST.
// If icons exist, we'll validate it.
func (lw *LinearWrapper) Validate() error {
	errors := make([]error, 0)
	if lw.VideoClicks != nil {
		if err := lw.VideoClicks.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	// No need to validate Icon which is for VAST 3.0 only.

	for _, tracking := range lw.Trackings {
		if err := tracking.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}
	if len(errors) > 0 {
		return vastbasic.ValidationError{Errs: errors}
	}
	return nil
}
