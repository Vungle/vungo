package vast3

import vastbasic "github.com/Vungle/vungo/vast/basic"

// Wrapper type represents a <Wrapper> element in VAST document that contains a single URI
// reference to a vendor ad server, a.k.a. third party ad server, where the actual VAST w
// content will be served. The vendor ad server may also provide additional wrapper which
// eventually resolves to the actual ad.
type Wrapper struct {
	AdSystem     *vastbasic.AdSystem     `xml:"AdSystem"`     // Required.
	VastAdTagURI vastbasic.TrimmedData   `xml:"VASTAdTagURI"` // Required.
	Impressions  []*vastbasic.Impression `xml:"Impression"`   // Required.
	Errors       []string                `xml:"Error,omitempty"`
	Creatives    []*Creative             `xml:"Creatives>Creative"` // Required.
	Extensions   []*vastbasic.Extension  `xml:"Extensions>Extension,omitempty"`
}

// Validate method validates the Wrapper according to the VAST.
// AdSystem, VastAdTagURI, and Impressions are required.
// Creatives are optional, if it exists, we'll also validate it.
func (w *Wrapper) Validate() error {
	errors := make([]error, 0)
	if len(w.VastAdTagURI) == 0 {
		errors = append(errors, vastbasic.ErrWrapperMissVastAdTagURI)
	}

	if len(w.Impressions) == 0 {
		errors = append(errors, vastbasic.ErrWrapperMissImpressions)
	}

	for _, c := range w.Creatives {
		if err := c.Validate(); err != nil {
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
