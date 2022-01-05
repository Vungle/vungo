package vastelement

// ViewableImpressionResource type is used to provide tracking for both the Inline and any Wrappers. Vast 4.0.
type ViewableImpressionResource struct {
	ID  string      `xml:"id,attr,omitempty"`
	URI TrimmedData `xml:",cdata"`
}

// Validate method validate ViewableImpressionResource vast element
func (viewableImpressionResource *ViewableImpressionResource) Validate(version Version) error {
	return nil
}
