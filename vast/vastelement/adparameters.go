package vastelement

// AdParameters type represents an <AdParameters> element that defines arbitrary ad parameters.
type AdParameters struct {
	Parameters TrimmedData `xml:",cdata"`

	IsXMLEncoded bool `xml:"xmlEncoded,attr,omitempty"` // Vast 3.0
}
