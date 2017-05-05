package vast

// Ad type represent an <Ad> child element in a VAST document. An <Ad> element usually specifies
// creatives, prices, delivery method, targeting, etc.
//
// Each <Ad> contains either EXACTLY ONE <InLine> element or <Wrapper> element (but never both).
type Ad struct {
	Id       string `xml:"id,attr,omitempty"`       // Id of the ad, defined by ad server. Required in VAST2.0.
	Sequence int    `xml:"sequence,attr,omitempty"` // Sequence number in which an ad should play. VAST3.0.

	InLine  *InLine  `xml:",omitempty"`
	Wrapper *Wrapper `xml:",omitempty"`
}

// Validate methods validate the Ad element according to the VAST.
// Each <Ad> contains either EXACTLY ONE <InLine> element or <Wrapper> element (but never both).
func (ad *Ad) Validate() error {
	errors := make([]error, 0)
	if ad.InLine != nil {
		if ad.Wrapper != nil {
			errors = append(errors, ErrAdType)
		}
		if err := ad.InLine.Validate(); err != nil {
			ve, ok := err.(ValidationError)
			if ok {
				errors = append(errors, ve.Errs...)
			}
		}
	} else {
		if ad.Wrapper == nil {
			errors = append(errors, ErrAdType)
		}
		if err := ad.Wrapper.Validate(); err != nil {
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
