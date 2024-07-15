package vastelement

// Linear type represents an <Linear> element within a <InLine> element.
type Linear struct {
	ID string `xml:"id,attr,omitempty"` // ID of the Linear defined by the ad server.
	// SkipOffset is time delay at which ad becomes skippable;
	// if absent, the ad is unskippable. VAST3.0.
	SkipOffset    *Offset `xml:"skipoffset,attr,omitempty"`
	IsAlternative *bool   `xml:"alternative,attr,omitempty"`
	// IsLooping controls whether the video loops after it finishes playing.
	// This is currently supported for accelerate only.
	IsLooping    *bool         `xml:"looping,attr,omitempty"`
	Duration     Duration      `xml:"Duration"`               // Required in InLine, optional in Wrapper
	AdParameters *AdParameters `xml:"AdParameters,omitempty"` // Type changes from string to struct in VAST3.0.
	Trackings    []*Tracking   `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks  *VideoClicks  `xml:"VideoClicks,omitempty"`
	MediaFiles   []*MediaFile  `xml:"MediaFiles>MediaFile,omitempty"`

	Extensions []*Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
	Icons      []*Icon      `xml:"Icons>Icon"`                                     // VAST3.0.
}

// Validate methods validate the Linear element according to the VAST.
// Duration are required in Linear but not required in Wrapper.
// According to VAST2.0 protocol, MediaFile is not required.
// Icons are optional, but if contained, we'll also validate it.
func (linear *Linear) Validate(version Version) error {
	errors := make([]error, 0)

	if linear.VideoClicks != nil {
		if err := linear.VideoClicks.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if linear.SkipOffset != nil {
		if err := linear.SkipOffset.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	// No need to validate Icon which is for VAST 3.0 only.

	for _, tracking := range linear.Trackings {
		if err := tracking.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if version == Version3 {
		for _, icon := range linear.Icons {
			if err := icon.Validate(version); err != nil {
				ve, ok := err.(ValidationError)
				if ok {
					errors = append(errors, ve.Errs...)
				}
			}
		}
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
