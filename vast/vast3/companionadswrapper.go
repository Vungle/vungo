package vast3

import vastbasic "github.com/Vungle/vungo/vast/basic"

// CompanionAdsWrapper type represents a <CompanionAds> element within a <Wrapper> element that
// contains companion creatives in a wrapper.
type CompanionAdsWrapper struct {
	Required   vastbasic.Mode      `xml:"required,attr,omitempty"` // VAST3.0.
	Companions []*CompanionWrapper `xml:"Companion,omitempty"`
}

// Validate method validates the CompanionAdsWrapper according to the VAST.
func (cw *CompanionAdsWrapper) Validate() error {
	errors := make([]error, 0)
	if len(cw.Companions) == 0 {
		errors = append(errors, vastbasic.ErrCompanionAdsWrapperMissCompanions)
	}

	// No need to validate Required which is for VAST 3.0 only.

	for _, c := range cw.Companions {
		if err := c.Validate(); err != nil {
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
