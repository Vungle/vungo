package vast

import "github.com/Vungle/vungo/vast/defaults"

// Linear type represents an <Linear> element within a <InLine> element.
type Linear struct {
	// SkipOffset is time delay at which ad becomes skippable;
	// if absent, the ad is unskippable. VAST3.0.
	SkipOffset *Offset `xml:"skipoffset,attr,omitempty"`

	Duration     Duration      `xml:"Duration"`               // Required.
	AdParameters *AdParameters `xml:"AdParameters,omitempty"` // Just string in VAST2.0.
	Icons        []*Icon       `xml:"Icons>Icon"`             // VAST3.0.
	Trackings    []*Tracking   `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks  *VideoClicks  `xml:"VideoClicks,omitempty"`
	MediaFiles   []*MediaFile  `xml:"MediaFiles>MediaFile,omitempty"`
	Extensions   []*Extension  `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}

// Validate methods validate the Linear element according to the VAST.
// MediaFiles and Duration are required.
// Icons are optional, but if contained, we'll also validate it.
func (linear *Linear) Validate() error {
	errors := make([]error, 0)
	if len(linear.MediaFiles) == 0 {
		errors = append(errors, ErrLinearMissMediaFiles) // Can be zero?
	}

	if err := linear.Duration.Validate(); err != nil {
		return err
	}

	if linear.Duration > Duration(defaults.MAX_VIDEO_DURATION) {
		errors = append(errors, ErrVideoDurationTooLong)
	}

	if linear.Duration < Duration(defaults.MIN_VIDEO_DURATION) {
		errors = append(errors, ErrVideoDurationTooShort)
	}

	if linear.VideoClicks != nil {
		if err := linear.VideoClicks.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if linear.SkipOffset != nil {
		if err := linear.SkipOffset.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if linear.AdParameters != nil {
		if err := linear.AdParameters.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	var err error
	var validMedia *MediaFile
	for _, mediaFile := range linear.MediaFiles {
		err = mediaFile.Validate()
		if err == nil {
			validMedia = mediaFile
			break
		}
	}
	if err != nil {
		// TODO(tiagozortea): The error being returned is just the last media error, we should
		// apply the proper error generalizations.
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	} else {
		linear.MediaFiles = []*MediaFile{validMedia}
	}

	for _, icon := range linear.Icons {
		if err := icon.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, tracking := range linear.Trackings {
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
