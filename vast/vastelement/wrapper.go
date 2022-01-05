package vastelement

// Wrapper type represents a <Wrapper> element in VAST document that contains a single URI
// reference to a vendor ad server, a.k.a. third party ad server, where the actual VAST w
// content will be served. The vendor ad server may also provide additional wrapper which
// eventually resolves to the actual ad.
type Wrapper struct {
	AdSystem     *AdSystem     `xml:"AdSystem"`     // Required.
	VastAdTagURI TrimmedData   `xml:"VASTAdTagURI"` // Required.
	Impressions  []*Impression `xml:"Impression"`   // Required.
	Errors       []string      `xml:"Error,omitempty"`
	Creatives    []*Creative   `xml:"Creatives>Creative"` // Required.
	Extensions   []*Extension  `xml:"Extensions>Extension,omitempty"`

	AllowMultipleAds         bool                `xml:"allowMultipleAds,attr,omitempty"`         // Vast 4.0
	FallbackOnNoAd           bool                `xml:"fallbackOnNoAd,attr,omitempty"`           // Vast 4.0
	FollowAdditionalWrappers bool                `xml:"followAdditionalWrappers,attr,omitempty"` // Vast 4.0
	Pricing                  *Pricing            `xml:"Pricing,omitempty"`                       // Vast 4.0
	ViewableImpression       *ViewableImpression `xml:"ViewableImpression,omitempty"`            // Vast 4.0
	AdVerification           []*Verification     `xml:"AdVerifications>Verification,omitempty"`  // Vast 4.0
}

// Validate method validates the Wrapper according to the VAST.
// AdSystem, VastAdTagURI, and Impressions are required.
// Creatives are optional, if it exists, we'll also Validate it.
func (wrapper *Wrapper) Validate(version Version) error {
	errors := make([]error, 0)
	if len(wrapper.VastAdTagURI) == 0 {
		errors = append(errors, ErrWrapperMissVastAdTagURI)
	}

	if len(wrapper.Impressions) == 0 {
		errors = append(errors, ErrWrapperMissImpressions)
	}

	for _, c := range wrapper.Creatives {
		if err := c.Validate(version); err != nil {
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
