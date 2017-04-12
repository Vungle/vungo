package vast

// Wrapper type represents a <Wrapper> element in VAST document that contains a single URI
// reference to a vendor ad server, a.k.a. third party ad server, where the actual VAST inline
// content will be served. The vendor ad server may also provide additional wrapper which
// eventually resolves to the actual ad.
type Wrapper struct {
	AdSystem     *AdSystem          `xml:"AdSystem"`           // Required.
	VastAdTagUri string             `xml:"VASTAdTagURI"`       // Required.
	Impressions  []*Impression      `xml:"Impression"`         // Required. should be a URI slice in VAST2.0.
	Errors       []string           `xml:"Error,omitempty"`    // Error should be an URI instead of a slice in VAST2.0.
	Creatives    []*CreativeWrapper `xml:"Creatives>Creative"` // Required.
	Extensions   []*Extension       `xml:"Extensions>Extension,omitempty"`
}

// Validate method validates the Wrapper according to the VAST.
// AdSystem, VastAdTagUri, and Impressions are required.
// Creatives are optional, if it exists, we'll also validate it.
func (w *Wrapper) Validate() error {

	if err := w.AdSystem.Validate(); err != nil {
		return err
	}

	if len(w.VastAdTagUri) == 0 {
		return ErrWrapperMissVastAdTagUri
	}

	if len(w.Impressions) == 0 {
		return ErrWrapperMissImpressions
	}

	for _, impression := range w.Impressions {
		if err := impression.Validate(); err != nil {
			return err
		}
	}

	for _, c := range w.Creatives {
		if err := c.Validate(); err != nil {
			return err
		}
	}

	return nil
}
