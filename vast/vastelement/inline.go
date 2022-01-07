package vastelement

import "github.com/Vungle/vungo/vast/defaults"

// InLine type represents an <InLine> ad element contained with <Ad> element in a VAST document.
//
// The <InLine> element contains all necessary assets, URI's, creatives, etc, to display the ad.
type InLine struct {
	AdSystem    *AdSystem     `xml:"AdSystem"`           // Required.
	AdTitle     *string       `xml:"AdTitle"`            // Required.
	Impressions []*Impression `xml:"Impression"`         // Required.
	Creatives   []*Creative   `xml:"Creatives>Creative"` // Creatives node is required.
	Description string        `xml:"Description,omitempty"`
	Survey      *Survey       `xml:"Survey,omitempty"`
	Errors      []string      `xml:"Error,omitempty"`
	Extensions  []*Extension  `xml:"Extensions>Extension,omitempty"`

	Advertiser string   `xml:"Advertiser,omitempty"` // Vast 3.0
	Pricing    *Pricing `xml:"Pricing,omitempty"`    // Vast 3.0

	Category           []*Category         `xml:"Category,omitempty"`                     // Vast 4.0
	ViewableImpression *ViewableImpression `xml:"ViewableImpression,omitempty"`           // Vast 4.0
	AdVerification     []*Verification     `xml:"AdVerifications>Verification,omitempty"` // Vast 4.0
}

// Validate methods validate the Inline element according to the VAST.
// AdSystem, AdTitle, Impression, Creatives are required.
func (inline *InLine) Validate(version Version) error {
	errors := make([]error, 0)
	if inline.AdTitle == nil {
		errors = append(errors, ErrInlineMissAdTitle)
	}

	if len(inline.Creatives) == 0 {
		errors = append(errors, ErrInlineMissCreatives)
	}

	// No need to validate Pricing which is for VAST 3.0 only.

	if len(inline.Impressions) == 0 {
		errors = append(errors, ErrInlineMissImpressions)
	}

	for _, creative := range inline.Creatives {
		if err := creative.Validate(version); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
		if creative.Linear != nil {
			if err := creative.Linear.Duration.Validate(version); err != nil {
				return err
			}

			if creative.Linear.Duration > Duration(defaults.MaxVideoDuration) {
				errors = append(errors, ErrVideoDurationTooLong)
			}

			if creative.Linear.Duration < Duration(defaults.MinVideoDuration) {
				errors = append(errors, ErrVideoDurationTooShort)
			}

			if len(creative.Linear.MediaFiles) == 0 {
				errors = append(errors, ErrLinearMissMediaFiles)
			}

			var validMediaFiles []*MediaFile
			var err error

			var hasMimeTypeErr bool
			noneMimeTypeErrors := make([]error, 0)

			for _, mediaFile := range creative.Linear.MediaFiles {
				err = mediaFile.Validate(version)
				if err == nil {
					validMediaFiles = append(validMediaFiles, mediaFile)
					//break
				} else {
					ve, ok := err.(ValidationError)
					if ok {
						// Merge all errors which are not mime type errors.
						if ve.Errs[0] != ErrMediaFileUnsupportedMimeType {
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
					errors = append(errors, ErrMediaFileUnsupportedMimeType)
				}
			}
		}
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}
