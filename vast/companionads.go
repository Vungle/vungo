package vast

// CompanionAds type represents a <CompanionAds> element that contains companion creatives.
type CompanionAds struct {
	Required   Mode         `xml:"required,attr,omitempty"` // VAST3.0.
	Companions []*Companion `xml:"Companion,omitempty"`
}

// Validate methods validate the CompanionAds element according to the VAST.
func (companionAds *CompanionAds) Validate() error {
	errors := make([]error, 0)

	// No need to validate Required which is for VAST 3.0 only.

	for _, companion := range companionAds.Companions {
		if err := companion.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
