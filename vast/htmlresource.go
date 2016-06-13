package vast

// HtmlResource type represents an <HtmlResource> element that contains HTML data.
type HtmlResource struct {
	IsXmlEncoded bool        `xml:"xmlEncoded,attr,omitempty"`
	Html         TrimmedData `xml:",cdata"`
}

// Validate method validates the HtmlResource according to the VAST.
// Html is required.
func (r *HtmlResource) Validate() error {

	if len(r.Html) == 0 {
		return ErrHtmlResourceMissHtml
	}

	return nil
}
