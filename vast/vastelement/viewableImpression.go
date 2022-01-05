package vastelement

// ViewableImpression type is used to provide tracking for both the Inline and any Wrappers. Vast 4.0.
type ViewableImpression struct {
	ID               string      `xml:"id,attr,omitempty"`
	Viewable         TrimmedData `xml:"Viewable,omitempty"`
	NotViewable      TrimmedData `xml:"NotViewable,omitempty"`
	ViewUndetermined TrimmedData `xml:"ViewUndetermined,omitempty"`
}

// Validate method validate ViewableImpression vast element
func (viewableImpression *ViewableImpression) Validate(version Version) error {
	return nil
}
