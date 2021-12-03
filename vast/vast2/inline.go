package vast2

import "github.com/Vungle/vungo/vast/basic"

// InLine type represents an <InLine> ad element contained with <Ad> element in a VAST document.
//
// The <InLine> element contains all necessary assets, URI's, creatives, etc, to display the ad.
type InLine struct {
	AdSystem    *vastbasic.AdSystem     `xml:"AdSystem"`           // Required.
	AdTitle     *string                 `xml:"AdTitle"`            // Required.
	Impressions []*vastbasic.Impression `xml:"Impression"`         // Required.
	Creatives   []*Creative             `xml:"Creatives>Creative"` // Creatives node is required.
	Description string                  `xml:"Description,omitempty"`
	Advertiser  string                  `xml:"Advertiser,omitempty"` // VAST3.0.
	SurveyURL   string                  `xml:"Survey,omitempty"`
	Errors      []string                `xml:"Error,omitempty"`
	Extensions  []*vastbasic.Extension  `xml:"Extensions>Extension,omitempty"`
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
	}

	if len(errors) > 0 {
		return vastbasic.ValidationError{Errs: errors}
	}
	return nil
}
