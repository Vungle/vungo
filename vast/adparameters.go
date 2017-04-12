package vast

// AdParameters type represents an <AdParameters> element that defines arbitrary ad parameters.
type AdParameters struct {
	Parameters TrimmedData `xml:",cdata"`

	IsXmlEncoded bool `xml:"xmlEncoded,attr,omitempty"` // VAST3.0.
}

// Validate method validates the AdParameters according to the VAST.
// Parameters are required.
func (p *AdParameters) Validate() error {

	if len(p.Parameters) == 0 {
		return ErrAdParametersMissParameters
	}

	return nil
}
