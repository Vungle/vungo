package vast

// HtmlResource type represents an <HtmlResource> element that contains HTML data.
type HtmlResource struct {
	Html         TrimmedData `xml:",cdata"`
	IsXmlEncoded bool        `xml:"xmlEncoded,attr,omitempty"`
}

// Validate method validates the HtmlResource according to the VAST.
// Html is required.
func (r *HtmlResource) Validate() error {

	if len(r.Html) == 0 {
		return ErrHtmlResourceMissHtml
	}

	return nil
}
