package vastelement

// Verification element is used to contain the JavaScript or Flash code used to collect data. Vast 4.0
type Verification struct {
	Vendor             string                        `xml:"vendor,attr,omitempty"`        // Vast 4.0
	JavaScriptResource []*JavaScriptResource         `xml:"JavaScriptResource,omitempty"` // Vast 4.0
	FlashResource      []*FlashResource              `xml:"FlashResource,omitempty"`      // Vast 4.0
	ViewableImpression []*ViewableImpressionResource `xml:"ViewableImpression,omitempty"` // Vast 4.0
}

// Validate method validate Survey vast element
func (verification *Verification) Validate(version Version) error {
	return nil
}
