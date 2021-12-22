package entity

import (
	vastbasic "github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/defaults"
)

// InLine type represents an <InLine> ad element contained with <Ad> element in a VAST document.
//
// The <InLine> element contains all necessary assets, URI's, creatives, etc, to display the ad.
type InLine struct {
	AdSystem    *vastbasic.AdSystem     `xml:"AdSystem"`           // Required.
	AdTitle     *string                 `xml:"AdTitle"`            // Required.
	Impressions []*vastbasic.Impression `xml:"Impression"`         // Required.
	Creatives   []*Creative             `xml:"Creatives>Creative"` // Creatives node is required.
	Description string                  `xml:"Description,omitempty"`
	SurveyURL   string                  `xml:"Survey,omitempty"`
	Errors      []string                `xml:"Error,omitempty"`
	Extensions  []*vastbasic.Extension  `xml:"Extensions>Extension,omitempty"`

	Advertiser string             `xml:"Advertiser,omitempty"` // VAST3.0.
	Pricing    *vastbasic.Pricing `xml:"Pricing,omitempty"`    // VAST3.0.
}

// Validate methods validate the Inline element according to the VAST.
// AdSystem, AdTitle, Impression, Creatives are required.
func (inline *InLine) Validate(version vastbasic.Version) error {
	errors := make([]error, 0)
	if inline.AdTitle == nil {
		errors = append(errors, vastbasic.ErrInlineMissAdTitle)
	}

	if len(inline.Creatives) == 0 {
		errors = append(errors, vastbasic.ErrInlineMissCreatives)
	}

	// No need to validate Pricing which is for VAST 3.0 only.

	if len(inline.Impressions) == 0 {
		errors = append(errors, vastbasic.ErrInlineMissImpressions)
	}

	for _, creative := range inline.Creatives {
		if err := creative.Validate(version); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
		if creative.Linear != nil {
			if err := creative.Linear.Duration.Validate(version); err != nil {
				return err
			}

			if creative.Linear.Duration > vastbasic.Duration(defaults.MaxVideoDuration) {
				errors = append(errors, vastbasic.ErrVideoDurationTooLong)
			}

			if creative.Linear.Duration < vastbasic.Duration(defaults.MinVideoDuration) {
				errors = append(errors, vastbasic.ErrVideoDurationTooShort)
			}

			if len(creative.Linear.MediaFiles) == 0 {
				errors = append(errors, vastbasic.ErrLinearMissMediaFiles)
			}

			var validMediaFiles []*vastbasic.MediaFile
			var err error

			var hasMimeTypeErr bool
			noneMimeTypeErrors := make([]error, 0)

			for _, mediaFile := range creative.Linear.MediaFiles {
				err = mediaFile.Validate(version)
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
				creative.Linear.MediaFiles = validMediaFiles
			} else {
				if len(noneMimeTypeErrors) > 0 {
					errors = append(errors, noneMimeTypeErrors...)
				}
				if hasMimeTypeErr {
					errors = append(errors, vastbasic.ErrMediaFileUnsupportedMimeType)
				}
			}
		}
	}

	if len(errors) > 0 {
		return vastbasic.ValidationError{Errs: errors}
	}
	return nil
}
