package vast2

// VideoClick type represents a complex type in VAST document that defines a click URL for a linear
// creative.
type VideoClick struct {
	ID  string      `xml:"id,attr,omitempty"`
	URI TrimmedData `xml:",cdata"`
}

// Validate method validates VideoClick according to the VAST.
func (click *VideoClick) Validate() error {
	return nil
}
