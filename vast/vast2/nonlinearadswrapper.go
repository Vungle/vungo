package vast2

import "github.com/Vungle/vungo/vast/basic"

// NonLinearAdsWrapper type represents a <NonLinearAds> element within a <Wrapper> element that
// contains non-linear creatives in a wrapper.
type NonLinearAdsWrapper struct {
	Trackings []*vastbasic.Tracking `xml:"TrackingEvents>Tracking,omitempty"`

	// TODO(@DevKai): NonLinears has to be changed to NonLinearWrapper in VAST3.0.
	NonLinears []*NonLinear `xml:"NonLinear,omitempty"`
}

// Validate method validates the NonLinearAdsWrapper according to the VAST.
func (nw *NonLinearAdsWrapper) Validate() error {
	errors := make([]error, 0)
	for _, tracking := range nw.Trackings {
		if err := tracking.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, nonLinear := range nw.NonLinears {
		if err := nonLinear.Validate(); err != nil {
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
