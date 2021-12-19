package vastbasic

// VideoClick type represents a complex type in VAST document that defines a click URL for a linear
// creative.
type VideoClick struct {
	ID  string      `xml:"id,attr,omitempty"`
	URI TrimmedData `xml:",cdata"`
}
