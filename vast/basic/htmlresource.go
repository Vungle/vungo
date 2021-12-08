package vastbasic

// HTMLResource type represents an <HTMLResource> element that contains HTML data.
type HTMLResource struct {
	HTML         TrimmedData `xml:",cdata"`
	IsXMLEncoded bool        `xml:"xmlEncoded,attr,omitempty"`
}
