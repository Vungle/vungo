package vast2

import (
	"github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/defaults"
)

// Linear type represents an <Linear> element within a <InLine> element.
type Linear struct {
	// SkipOffset is time delay at which ad becomes skippable;
	Duration     vastbasic.Duration     `xml:"Duration"`               // Required.
	AdParameters *string                `xml:"AdParameters,omitempty"` // Just string in VAST2.0.
	Trackings    []*vastbasic.Tracking  `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks  *vastbasic.VideoClicks `xml:"VideoClicks,omitempty"`
	MediaFiles   []*vastbasic.MediaFile `xml:"MediaFiles>MediaFile,omitempty"`
}

// Validate methods validate the Linear element according to the VAST.
// MediaFiles and Duration are required.
// Icons are optional, but if contained, we'll also validate it.
func (linear *Linear) Validate() error {
	errors := make([]error, 0)
	if len(linear.MediaFiles) == 0 {
		errors = append(errors, vastbasic.ErrLinearMissMediaFiles) // Can be zero?
	}

	if err := linear.Duration.Validate(); err != nil {
		return err
	}

	if linear.Duration > vastbasic.Duration(defaults.MaxVideoDuration) {
		errors = append(errors, vastbasic.ErrVideoDurationTooLong)
	}

	if linear.Duration < vastbasic.Duration(defaults.MinVideoDuration) {
		errors = append(errors, vastbasic.ErrVideoDurationTooShort)
	}

	if linear.VideoClicks != nil {
		if err := linear.VideoClicks.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	var validMediaFiles []*vastbasic.MediaFile
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
