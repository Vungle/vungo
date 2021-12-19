package vastbasic

// Impression type represents a URI that directs the video player to a tracking resource file that
// the video player should request when the first frame of the ad is displayed.
type Impression struct {
	ID  string      `xml:"id,attr,omitempty"`
	URI TrimmedData `xml:",cdata"`
}
