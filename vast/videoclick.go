package vast

// VideoClick type represents a complex type in VAST document that defines a click URL for a linear
// creative.
type VideoClick struct {
	Id  string      `xml:"id,attr,omitempty"`
	Uri TrimmedData `xml:",cdata"`
}

// Validate method validates VideoClick according to the VAST.
// Uri is required.
func (click *VideoClick) Validate() error {

	if len(click.Uri) == 0 {
		return ErrVideoClickMissUri
	}

	return nil
}
