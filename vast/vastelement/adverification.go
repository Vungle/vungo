package vastelement

// Verification type represents an <Verification> element that contains arbitrary XML
// data provided by the platform to extend the VAST response.
type Verification struct {
	Vendor                 string                  `xml:"vendor,attr"`
	ViewableImpression     *ViewableImpression     `xml:"ViewableImpression,omitempty"` // VAST4.0
	JavaScriptResource     *JavaScriptResource     `xml:"JavaScriptResource,omitempty"`
	FlashResource          *FlashResource          `xml:"FlashResource,omitempty"`           // VAST4.0
	ExecutableResource     *ExecutableResource     `xml:"ExecutableResource,omitempty"`      // VAST4.1
	VerificationParameters *VerificationParameters `xml:"VerificationParameters,omitempty"`  // VAST4.1
	TrackingEvents         []*Tracking             `xml:"TrackingEvents>Tracking,omitempty"` // VAST4.1
}

// JavaScriptResource vast url for javascript res
type JavaScriptResource struct {
	APIFramework    string `xml:"apiFramework,attr,omitempty"`
	BrowserOptional string `xml:"browserOptional,attr,omitempty"`
	Value           string `xml:",cdata"`
}

// FlashResource is a container for the URI to the Flash file used to collect verification data.
type FlashResource struct {
	APIFramework string `xml:"apiFramework,attr,omitempty"`
	Value        string `xml:",cdata"`
}

// ExecutableResource is a container for the URI to the executable file used to collect verification data.
type ExecutableResource struct {
	APIFramework string `xml:"apiFramework,attr,omitempty"`
	Value        string `xml:",cdata"`
}

// VerificationParameters vast element
type VerificationParameters struct {
	Value TrimmedData `xml:",cdata"`
}

// ViewableImpression type represents an <ViewableImpression> element that contains arbitrary XML
// data provided by the platform to extend the VAST response.
type ViewableImpression struct {
	ID   string      `xml:"id,attr"`
	Data TrimmedData `xml:",cdata"`
}
