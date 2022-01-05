package vastelement

// JavaScriptResource is a container for the URI to the JavaScript file used to collect verification data.
type JavaScriptResource struct {
	APIFramework string      `xml:"apiFramework,attr,omitempty"`
	URI          TrimmedData `xml:",cdata"`
}

// Validate method validates JavaScriptResource according to the VAST.
func (javaScriptResource *JavaScriptResource) Validate(version Version) error {
	return nil
}
