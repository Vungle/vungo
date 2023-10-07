package vastelement

// Ad type represent an <Ad> child element in a VAST document. An <Ad> element usually specifies
// creatives, prices, delivery method, targeting, etc.
//
// Each <Ad> contains either EXACTLY ONE <InLine> element or <Wrapper> element (but never both).
type Ad struct {
	ID      string   `xml:"id,attr,omitempty"` // ID of the ad, defined by ad server. Required in VAST2.0.
	InLine  *InLine  `xml:",omitempty"`
	Wrapper *Wrapper `xml:",omitempty"`

	Sequence int `xml:"sequence,attr,omitempty"` // Sequence number in which an ad should play. VAST3.0.
}

// Validate methods validate the Ad element according to the VAST.
// Each <Ad> contains either EXACTLY ONE <InLine> element or <Wrapper> element (but never both).
func (ad *Ad) Validate(version Version,needFilterURI bool) error {
	errors := make([]error, 0)

	if ad.InLine != nil && ad.Wrapper != nil {
		errors = append(errors, ErrAdType)
		return ValidationError{Errs: errors}
	}

	if ad.InLine != nil {
		errors = append(errors, ad.validateAdInline(version,needFilterURI)...)
	} else if ad.Wrapper != nil {
		errors = append(errors, ad.validateAdWrapper(version)...)
	} else {
		errors = append(errors, ErrAdType)
	}

	if len(errors) > 0 {
		return ValidationError{Errs: errors}
	}
	return nil
}

// validateAdInline method validate Inline in an Ad vast element
func (ad *Ad) validateAdInline(version Version,needFilterURI bool) []error {
	errors := make([]error, 0)
	if err := ad.InLine.Validate(version,needFilterURI bool); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	return errors
}

// validateAdWrapper method validate Wrapper in an Ad vast element
func (ad *Ad) validateAdWrapper(version Version) []error {
	errors := make([]error, 0)
	if err := ad.Wrapper.Validate(version); err != nil {
		ve, ok := err.(ValidationError)
		if ok {
			errors = append(errors, ve.Errs...)
		}
	}

	return errors
}
