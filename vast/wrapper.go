package vast

// Wrapper type represents a <Wrapper> element in VAST document that contains a single URI
// reference to a vendor ad server, a.k.a. third party ad server, where the actual VAST w
// content will be served. The vendor ad server may also provide additional wrapper which
// eventually resolves to the actual ad.
type Wrapper struct {
	AdSystem     *AdSystem          `xml:"AdSystem"`     // Required.
	VastAdTagUri string             `xml:"VASTAdTagURI"` // Required.
	Impressions  []*Impression      `xml:"Impression"`   // Required.
	Errors       []string           `xml:"Error,omitempty"`
	Creatives    []*CreativeWrapper `xml:"Creatives>Creative"` // Required.
	Extensions   []*Extension       `xml:"Extensions>Extension,omitempty"`
}

// Validate method validates the Wrapper according to the VAST.
// AdSystem, VastAdTagUri, and Impressions are required.
// Creatives are optional, if it exists, we'll also validate it.
func (w *Wrapper) Validate() error {
	errors := make([]error, 0)
	if w.AdSystem == nil {
		errors = append(errors, ErrWrapperMissAdSystem)
	}

	if len(w.VastAdTagUri) == 0 {
		errors = append(errors, ErrWrapperMissVastAdTagUri)
	}

	if len(w.Impressions) == 0 {
		errors = append(errors, ErrWrapperMissImpressions)
	}

	for _, c := range w.Creatives {
		if err := c.Validate(); err != nil {
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
