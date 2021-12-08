package vast3

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// Linear type represents an <Linear> element within a <InLine> element.
type Linear struct {
	// SkipOffset is time delay at which ad becomes skippable;
	// if absent, the ad is unskippable. VAST3.0.
	SkipOffset *vastbasic.Offset `xml:"skipoffset,attr,omitempty"`

	Duration     vastbasic.Duration     `xml:"Duration"`               // Required in InLine, optional in Wrapper
	AdParameters *AdParameters          `xml:"AdParameters,omitempty"` // Type changes from string to struct in VAST3.0.
	Icons        []*vastbasic.Icon      `xml:"Icons>Icon"`             // VAST3.0.
	Trackings    []*Tracking            `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks  *vastbasic.VideoClicks `xml:"VideoClicks,omitempty"`
	MediaFiles   []*MediaFile           `xml:"MediaFiles>MediaFile,omitempty"`
	Extensions   []*vastbasic.Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
}

// Validate methods validate the Linear element according to the VAST.
// Duration are required in Linear but not required in Wrapper.
// According to VAST2.0 protocol, MediaFile is not required.
// Icons are optional, but if contained, we'll also validate it.
func (linear *Linear) Validate() error {
	errors := make([]error, 0)

	if len(linear.MediaFiles) == 0 {
		errors = append(errors, vastbasic.ErrLinearMissMediaFiles) // Can be zero?
	}

	if linear.VideoClicks != nil {
		if err := linear.VideoClicks.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if linear.SkipOffset != nil {
		if err := linear.SkipOffset.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	var validMediaFiles []*MediaFile
	var err error

	var hasMimeTypeErr bool
	noneMimeTypeErrors := make([]error, 0)
	for _, mediaFile := range linear.MediaFiles {
		err = mediaFile.Validate()
		if err == nil {
			validMediaFiles = append(validMediaFiles, mediaFile)
			//break
		} else {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				// Merge all errors which are not mime type errors.
				if ve.Errs[0] != vastbasic.ErrMediaFileUnsupportedMimeType {
					noneMimeTypeErrors = append(noneMimeTypeErrors, ve.Errs...)
				} else {
					hasMimeTypeErr = true
				}
			}
		}
	}

	if len(validMediaFiles) > 0 {
		linear.MediaFiles = validMediaFiles
	} else {
		if len(noneMimeTypeErrors) > 0 {
			errors = append(errors, noneMimeTypeErrors...)
		}
		if hasMimeTypeErr {
			errors = append(errors, vastbasic.ErrMediaFileUnsupportedMimeType)
		}
	}

	// No need to validate Icon which is for VAST 3.0 only.

	for _, tracking := range linear.Trackings {
		if err := tracking.Validate(); err != nil {
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
