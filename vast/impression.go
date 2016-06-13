package vast

// Impression type represents a URI that directs the video player to a tracking resource file that
// the video player should request when the first frame of the ad is displayed.
type Impression struct {
	Id string `xml:"id,attr,omitempty"`

	Uri TrimmedData `xml:",cdata"`
}

// Validate methods validate the Impression element according to the VAST.
// Uri is required.
func (impression *Impression) Validate() error {

	if len(impression.Uri) == 0 {
		return ErrImpressionMissUri
	}

	return nil
}
