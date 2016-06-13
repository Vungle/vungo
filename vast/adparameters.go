package vast

// AdParameters type represents an <AdParameters> element that defines arbitrary ad parameters.
type AdParameters struct {
	IsXmlEncoded bool `xml:"xmlEncoded,attr,omitempty"`

	Parameters TrimmedData `xml:",cdata"`
}

// Validate method validates the AdParameters according to the VAST.
// Parameters are required.
func (p *AdParameters) Validate() error {

	if len(p.Parameters) == 0 {
		return ErrAdParametersMissParameters
	}

	return nil
}
