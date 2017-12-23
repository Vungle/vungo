package vast

// HtmlResource type represents an <HtmlResource> element that contains HTML data.
type HtmlResource struct {
	Html         TrimmedData `xml:",cdata"`
	IsXmlEncoded bool        `xml:"xmlEncoded,attr,omitempty"`
}
