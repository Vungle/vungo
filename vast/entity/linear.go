package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
)

// Linear type represents an <Linear> element within a <InLine> element.
type Linear struct {
	// SkipOffset is time delay at which ad becomes skippable;
	// if absent, the ad is unskippable. VAST3.0.
	SkipOffset   *vastbasic.Offset       `xml:"skipoffset,attr,omitempty"`
	Duration     vastbasic.Duration      `xml:"Duration"`               // Required in InLine, optional in Wrapper
	AdParameters *vastbasic.AdParameters `xml:"AdParameters,omitempty"` // Type changes from string to struct in VAST3.0.
	Trackings    []*vastbasic.Tracking   `xml:"TrackingEvents>Tracking,omitempty"`
	VideoClicks  *vastbasic.VideoClicks  `xml:"VideoClicks,omitempty"`
	MediaFiles   []*vastbasic.MediaFile  `xml:"MediaFiles>MediaFile,omitempty"`

	Extensions []*vastbasic.Extension `xml:"CreativeExtensions>CreativeExtension,omitempty"` // VAST3.0.
	Icons      []*vastbasic.Icon      `xml:"Icons>Icon"`                                     // VAST3.0.
}

// Validate methods validate the Linear element according to the VAST.
// Duration are required in Linear but not required in Wrapper.
// According to VAST2.0 protocol, MediaFile is not required.
// Icons are optional, but if contained, we'll also validate it.
func (linear *Linear) Validate(version vastbasic.Version) error {
	errors := make([]error, 0)

	if linear.VideoClicks != nil {
		if err := linear.VideoClicks.Validate(version); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	if linear.SkipOffset != nil {
		if err := linear.SkipOffset.Validate(version); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	// No need to validate Icon which is for VAST 3.0 only.

	for _, tracking := range linear.Trackings {
		if err := tracking.Validate(version); err != nil {
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
