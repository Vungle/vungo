package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// NonLinearAds type represents a <NonLinearAds> element that contains non-linear creatives.
type NonLinearAds struct {
	Trackings  []*vastbasic.Tracking `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinears []*NonLinear          `xml:"NonLinear,omitempty"` // Required, at least one item.
}

// Validate method validates the NonLinearAds according to the VAST.
// NonLinears is required.
func (nonLinearAds *NonLinearAds) Validate(version vastbasic.Version) error {
	errors := make([]error, 0)
	if len(nonLinearAds.NonLinears) == 0 {
		errors = append(errors, vastbasic.ErrNonLinearAdsMissNonLinears)
	}

	for _, nonLinear := range nonLinearAds.NonLinears {
		if err := nonLinear.Validate(version); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, tracking := range nonLinearAds.Trackings {
		if err := tracking.Validate(version); err != nil {
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
