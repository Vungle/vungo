package vast

// Extension type represents an <Extension> element that contains arbitrary XML data provided by the
// platform to extend the VAST response.
type Extension struct {
	Data          []byte          `xml:",innerxml"`
	Verifications []*Verification `xml:"AdVerifications>Verification,omitempty"`
}

type Verification struct {
	Vendor             string              `xml:"vendor,attr"`
	ViewableImpression *ViewableImpression `xml:"ViewableImpression"`
}

type ViewableImpression struct {
	ID   string      `xml:"id,attr"`
	Data TrimmedData `xml:",cdata"`
}
