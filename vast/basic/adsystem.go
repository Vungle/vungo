package vastbasic

// AdSystem type represents information about the ad server that returned the ad.
type AdSystem struct {
	System  string `xml:",chardata"`
	Version string `xml:"version,attr,omitempty"`
}
