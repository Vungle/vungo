package vastelement

import "encoding/xml"

// Vast type represents the root <VAST> tag.
type Vast struct {
	XMLName xml.Name `xml:"VAST"`
	Version Version  `xml:"version,attr"`    // Required.
	Ads     []*Ad    `xml:"Ad"`              // Ad can be empty in VAST2.0.
	Errors  []string `xml:"Error,omitempty"` // One or more URI's, likely tracking pixels, to request in case of no ad.
}

// FindFirstInlineLinearCreative method inspects through all of its inline ads and finds the first
// linear creative within, or returns nil when found nothing.
func (v Vast) FindFirstInlineLinearCreative() *Linear {
	for _, ad := range v.Ads {
		if ad.InLine == nil {
			continue
		}

		for _, c := range ad.InLine.Creatives {
			if c.Linear != nil {
				return c.Linear
			}
		}
	}
	return nil
}

// FindFirstInlineCompanionAdsCreative method inspects through all of its inline ads and finds the first
// CompanionAds creative within, or returns nil when found nothing.
func (v Vast) FindFirstInlineCompanionAdsCreative() *CompanionAds {
	for _, ad := range v.Ads {
		if ad.InLine == nil {
			continue
		}

		for _, c := range ad.InLine.Creatives {
			if c.CompanionAds != nil {
				return c.CompanionAds
			}
		}
	}
	return nil
}

// MediaFileChecker is a custom filter, input is MediaFile and the output is an error.
type MediaFileChecker func(file *MediaFile) error

// Validate method validates the Vast element according to the VAST.
// Version, and Ads are required.
func (v *Vast) Validate(version Version, checker MediaFileChecker) error {
	errors := make([]error, 0)
	if err := v.Version.Validate(version); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	if len(v.Ads) == 0 {
		errors = append(errors, ErrVastMissAd)
	}

	for _, ad := range v.Ads {
		if err := ad.Validate(version, checker); err != nil {
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
