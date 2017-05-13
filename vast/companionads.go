package vast

// CompanionAds type represents a <CompanionAds> element that contains companion creatives.
type CompanionAds struct {
	Required   Mode         `xml:"required,attr,omitempty"` // VAST3.0.
	Companions []*Companion `xml:"Companion,omitempty"`
}

// Validate methods validate the CompanionAds element according to the VAST.
func (companionAds *CompanionAds) Validate() error {
	errors := make([]error, 0)
	if len(companionAds.Companions) == 0 {
		errors = append(errors, ErrCompanionAdsMissCompanions)
	}

	if len(companionAds.Required) != 0 {
		if err := companionAds.Required.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

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
