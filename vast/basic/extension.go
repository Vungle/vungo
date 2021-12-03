package vastbasic

// Extension type represents an <Extension> element that contains arbitrary XML data provided by the
// platform to extend the VAST response.
type Extension struct {
	Data []byte `xml:",innerxml"`
}
