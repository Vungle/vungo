package vast

// StaticResource type represents an <StaticResource> element that contains information about a
// static resource, such as URL of the image of SWF file.
type StaticResource struct {
	MimeType string      `xml:"creativeType,attr"` // Required.
	Uri      TrimmedData `xml:",cdata"`
}

// Validate method validates the StaticResource according to the VAST.
// Uri is required.
func (r *StaticResource) Validate() error {
	if len(r.Uri) == 0 {
		return ValidationError{Errs: []error{ErrStaticResourceMissUri}}
	}

	return nil
}
