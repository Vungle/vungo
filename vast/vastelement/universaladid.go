package vastelement

// UniversalAdId type provides a unique creative identifier that is maintained across systems. Vast 4.0.
type UniversalAdId struct {
	IdRegistry string      `xml:"idRegistry,attr"`
	IdValue    string      `xml:"idValue,attr"`
	AdId       TrimmedData `xml:",chardata"`
}

// Validate method validate UniversalAdId vast element
func (universalAdId *UniversalAdId) Validate(version Version) error {
	return nil
}
