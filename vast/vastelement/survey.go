package vastelement

// Survey type can be used to provide a URI to any resource file having to do with collecting survey data.
type Survey struct {
	Type      string      `xml:"type,attr,omitempty"` // Vast 4.0
	SurveyURL TrimmedData `xml:",chardata"`
}

// Validate method validate Survey vast element
func (survey *Survey) Validate(version Version) error {
	return nil
}
