package vastbasic

// StaticResource type represents an <StaticResource> element that contains information about a
// static resource, such as URL of the image of SWF file.
type StaticResource struct {
	MimeType string      `xml:"creativeType,attr"` // Required.
	URI      TrimmedData `xml:",cdata"`
}
