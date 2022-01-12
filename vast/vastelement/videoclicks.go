package vastelement

// VideoClicks type represents a <VideoClicks> element that contains various types of VideoClick's.
type VideoClicks struct {
	ClickThroughs  []*VideoClick `xml:"ClickThrough,omitempty"`
	ClickTrackings []*VideoClick `xml:"ClickTracking,omitempty"`
	CustomClicks   []*VideoClick `xml:"CustomClick,omitempty"`
}

// Validate method validates the VideoClicks according to the VAST.
// ClickThroughs element is required.
func (videoClicks *VideoClicks) Validate(version Version) error {
	errors := make([]error, 0)

	// TODO(@DevKai): VAST3.0 require ClickThroughs not be empty.

	for _, click := range videoClicks.ClickThroughs {
		if err := click.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, click := range videoClicks.ClickTrackings {
		if err := click.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, click := range videoClicks.CustomClicks {
		if err := click.Validate(version); err != nil {
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
