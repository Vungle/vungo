package vast

// CreativeWrapper type represents a <Creative> element within a <Wrapper> element that defines
// a wrapped Creative's parent trackers.
type CreativeWrapper struct {
	ID       string `xml:"id,attr,omitempty"`
	Sequence int    `xml:"sequence,attr,omitempty"`
	AdID     string `xml:"AdID,attr,omitempty"`

	Linear       *LinearWrapper       `xml:"Linear,omitempty"`
	CompanionAds *CompanionAds        `xml:"CompanionAds,omitempty"`
	NonLinearAds *NonLinearAdsWrapper `xml:"NonLinearAds,omitempty"`
}

// Validate method validates the CreativeWrapper.
// CreativeWrapper can contain up to one of LinearWrapper, CompanionAdsWrapper, and NonLinearAdsWrapper.
func (cw *CreativeWrapper) Validate() error {
	errors := make([]error, 0)
	// Empty creative is valid for wrapper creative
	if cw.Linear == nil && cw.CompanionAds == nil && cw.NonLinearAds == nil {
		return nil
	}
	if cw.Linear != nil {
		if cw.CompanionAds != nil || cw.NonLinearAds != nil {
			errors = append(errors, ErrCreativeWrapperType)
		}
		if err := cw.Linear.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	} else if cw.CompanionAds != nil {
		if cw.NonLinearAds != nil {
			errors = append(errors, ErrCreativeWrapperType)
		}
		if err := cw.CompanionAds.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	} else if cw.NonLinearAds == nil {
		errors = append(errors, ErrCreativeWrapperType)
	} else if err := cw.NonLinearAds.Validate(); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}
	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
