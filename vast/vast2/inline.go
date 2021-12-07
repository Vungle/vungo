package vast2

import (
	"github.com/Vungle/vungo/vast/basic"
	"github.com/Vungle/vungo/vast/defaults"
)

// InLine type represents an <InLine> ad element contained with <Ad> element in a VAST document.
//
// The <InLine> element contains all necessary assets, URI's, creatives, etc, to display the ad.
type InLine struct {
	AdSystem    *vastbasic.AdSystem     `xml:"AdSystem"`   // Required.
	AdTitle     *string                 `xml:"AdTitle"`    // Required.
	Impressions []*vastbasic.Impression `xml:"Impression"` // Required.
	Errors      []string                `xml:"Error,omitempty"`
	Creatives   []*Creative             `xml:"Creatives>Creative"` // Creatives node is required.
	Extensions  []*vastbasic.Extension  `xml:"Extensions>Extension,omitempty"`

	Description string `xml:"Description,omitempty"`
	SurveyURL   string `xml:"Survey,omitempty"`
}

// Validate methods validate the Inline element according to the VAST.
// AdSystem, AdTitle, Impression, Creatives are required.
func (inline *InLine) Validate() error {
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
		if err := creative.Validate(); err != nil {
			ve, ok := err.(vastbasic.ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
		// check linear duration
		if linear := creative.Linear; linear != nil {
			if err := creative.Linear.Duration.Validate(); err != nil {
				return err
			}
			if creative.Linear.Duration > vastbasic.Duration(defaults.MaxVideoDuration) {
				errors = append(errors, vastbasic.ErrVideoDurationTooLong)
			}
			if creative.Linear.Duration < vastbasic.Duration(defaults.MinVideoDuration) {
				errors = append(errors, vastbasic.ErrVideoDurationTooShort)
			}
		}
	}

	if len(errors) > 0 {
		return vastbasic.ValidationError{Errs: errors}
	}
	return nil
}
