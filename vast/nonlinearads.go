package vast

// NonLinearAds type represents a <NonLinearAds> element that contains non-linear creatives.
type NonLinearAds struct {
	Trackings  []*Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	NonLinears []*NonLinear `xml:"NonLinear,omitempty"` // Required, at least one item.
}

// Validate method validates the NonLinearAds according to the VAST.
// NonLinears is required.
func (nonLinearAds *NonLinearAds) Validate() error {
	errors := make([]error, 0)
	if len(nonLinearAds.NonLinears) == 0 {
		errors = append(errors, ErrNonLinearAdsMissNonLinears)
	}

	for _, nonLinear := range nonLinearAds.NonLinears {
		if err := nonLinear.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, tracking := range nonLinearAds.Trackings {
		if err := tracking.Validate(); err != nil {
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
