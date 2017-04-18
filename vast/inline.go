package vast

// InLine type represents an <InLine> ad element contained with <Ad> element in a VAST document.
//
// The <InLine> element contains all necessary assets, URI's, creatives, etc, to display the ad.
type InLine struct {
	AdSystem    *AdSystem     `xml:"AdSystem"`           // Required.
	AdTitle     string        `xml:"AdTitle"`            // Required.
	Impressions []*Impression `xml:"Impression"`         // Required.
	Creatives   []*Creative   `xml:"Creatives>Creative"` // Creatives node is required.
	Description string        `xml:"Description,omitempty"`
	Advertiser  string        `xml:"Advertiser,omitempty"` // VAST3.0.
	SurveyUrl   string        `xml:"Survey,omitempty"`
	Errors      []string      `xml:"Error,omitempty"`   // Error should be an URI in VAST2.0.
	Pricing     *Pricing      `xml:"Pricing,omitempty"` // VAST3.0.
	Extensions  []*Extension  `xml:"Extensions>Extension,omitempty"`
}

// Validate methods validate the Inline element according to the VAST.
// AdSystem, AdTitle, Impression, Creatives are required.
func (inline *InLine) Validate() error {
	errors := make([]error, 0)
	if err := inline.AdSystem.Validate(); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	if len(inline.AdTitle) == 0 {
		errors = append(errors, ErrInlineMissAdTitle)
	}

	if len(inline.Impressions) == 0 {
		errors = append(errors, ErrInlineMissImpressions)
	}

	if len(inline.Creatives) == 0 {
		errors = append(errors, ErrInlineMissCreatives)
	}

	if inline.Pricing != nil {
		if err := inline.Pricing.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, impression := range inline.Impressions {
		if err := impression.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	}

	for _, creative := range inline.Creatives {
		if err := creative.Validate(); err != nil {
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
