package vastelement

// FlashResource is a container for the URI to the Flash file used to collect verification data.
type FlashResource struct {
	APIFramework string      `xml:"apiFramework,attr,omitempty"`
	URI          TrimmedData `xml:",cdata"`
}

// Validate method validates FlashResource according to the VAST.
func (flashResource *FlashResource) Validate(version Version) error {
	return nil
}
