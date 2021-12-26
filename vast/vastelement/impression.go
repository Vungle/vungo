package vastelement

// Impression type represents a URI that directs the video player to a tracking resource file that
// the video player should request when the first frame of the ad is displayed.
type Impression struct {
	ID  string      `xml:"id,attr,omitempty"`
	URI TrimmedData `xml:",cdata"`
}

// Validate method validate Impression vast element
func (impression *Impression) Validate(version Version) error {
	if len(impression.URI) == 0 {
		return ValidationError{Errs: []error{ErrImpressionMissURI}}
	}
	return nil
}
