package vastelement

// StaticResource type represents an <StaticResource> element that contains information about a
// static resource, such as URL of the image of SWF file.
type StaticResource struct {
	MimeType string      `xml:"creativeType,attr"` // Required.
	URI      TrimmedData `xml:",cdata"`
}

// Validate method validates the StaticResource according to the VAST.
func (StaticResource *StaticResource) Validate(version Version) error {
	return nil
}
