package entity

import vastbasic "github.com/Vungle/vungo/vast/basic"

// CompanionAds type represents a <CompanionAds> element that contains companion creatives.
type CompanionAds struct {
	Companions []*Companion `xml:"Companion,omitempty"`

	Required vastbasic.Mode `xml:"required,attr,omitempty"` // VAST3.0.
}

// Validate methods validate the CompanionAds element according to the VAST.
func (companionAds *CompanionAds) Validate(version vastbasic.Version) error {
	errors := make([]error, 0)

	// No need to validate Required which is for VAST 3.0 only.

	for _, companion := range companionAds.Companions {
		if err := companion.Validate(version); err != nil {
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
